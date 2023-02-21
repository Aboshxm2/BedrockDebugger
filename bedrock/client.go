package bedrock

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"github.com/Aboshxm2/BedrockDebugger/auth"
	"github.com/Aboshxm2/BedrockDebugger/translate"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

type Client struct {
	conn        *minecraft.Conn
	OnClose     func(reason string)
	OnSpawn     func()
	OnMessage   func(message string)
	players     map[string]uint64
	pos         mgl32.Vec3
	sneaking    bool
	RainbowSkin bool
}

func NewClient(address string) (*Client, error) {
	conn, err := minecraft.Dialer{
		TokenSource: auth.TokenSource(),
	}.Dial("raknet", address)

	if err != nil {
		return nil, err
	}

	return &Client{conn: conn, players: map[string]uint64{}}, nil
}

func (c *Client) Connect() {
	if err := c.conn.DoSpawn(); err != nil {
		c.OnClose(err.Error())
		return
	}

	defer c.conn.Close()

	c.pos = c.conn.GameData().PlayerPosition

	c.OnSpawn()

	for {
		pk, err := c.conn.ReadPacket()

		if err != nil {
			if disconnect, ok := errors.Unwrap(err).(minecraft.DisconnectError); ok {
				c.OnClose(disconnect.Error())
			} else {
				c.OnClose(err.Error())
			}
			break
		}

		c.handel(pk)
	}
}

func (c *Client) handel(pk packet.Packet) {
	switch p := pk.(type) {
	case *packet.Text:
		if p.TextType == packet.TextTypeChat || p.TextType == packet.TextTypeRaw || p.TextType == packet.TextTypeTranslation {
			msg := text.Clean(p.Message)

			if p.TextType == packet.TextTypeTranslation {
				msg = translate.Translate(msg, p.Parameters)
			}

			c.OnMessage(msg)
		}
	case *packet.Disconnect:
		c.conn.Close()
		c.OnClose(p.Message)
	case *packet.AddPlayer:
		c.players[p.Username] = p.EntityRuntimeID
	case *packet.RemoveEntity:
		for player, id := range c.players {
			if id == p.EntityNetworkID {
				delete(c.players, player)
			}
		}
	case *packet.MovePlayer:
		if p.EntityRuntimeID == c.conn.GameData().EntityRuntimeID {
			c.pos = p.Position
		}
	}
}

func (c *Client) Chat(message string) {
	c.conn.WritePacket(&packet.Text{
		TextType:         packet.TextTypeChat,
		NeedsTranslation: false,
		SourceName:       "",
		Message:          message,
		Parameters:       []string{},
		XUID:             "",
		PlatformChatID:   "",
	})
}

func (c *Client) Attack(player string) {
	id, ok := c.players[player]
	if !ok {
		return
	}

	c.conn.WritePacket(&packet.InventoryTransaction{
		TransactionData: &protocol.UseItemOnEntityTransactionData{
			TargetEntityRuntimeID: id,
			ActionType:            protocol.UseItemOnEntityActionAttack,
		},
	})
}

func (c *Client) Jump() {
	c.conn.WritePacket(&packet.PlayerAuthInput{
		Position:  c.pos,
		InputData: packet.InputFlagStartJumping,
	})
}

func (c *Client) Sneak() {

	if c.sneaking {
		c.conn.WritePacket(&packet.PlayerAuthInput{
			Position:  c.pos,
			InputData: packet.InputFlagStartSneaking,
		})
	} else {
		c.conn.WritePacket(&packet.PlayerAuthInput{
			Position:  c.pos,
			InputData: packet.InputFlagStopSneaking,
		})
	}

	c.sneaking = !c.sneaking
}

func (c *Client) Move(direction string) {
	pos := c.pos

	switch direction {
	case "east":
		pos = pos.Add(mgl32.Vec3{0.9})
	case "west":
		pos = pos.Add(mgl32.Vec3{-0.9})
	case "north":
		pos = pos.Add(mgl32.Vec3{0, 0, 0.9})
	case "south":
		pos = pos.Add(mgl32.Vec3{0, 0, -0.9})
	case "up":
		pos = pos.Add(mgl32.Vec3{0, 0.9})
	case "down":
		pos = pos.Add(mgl32.Vec3{0, -0.9})
	}

	c.conn.WritePacket(&packet.PlayerAuthInput{
		Position: pos,
	})

	c.pos = pos
}

func (c *Client) Disconnect(reason string) {
	c.conn.Close()
	c.OnClose(reason)
}

func (c *Client) Rainbow() {
	go func() {
		c.RainbowSkin = true
		i := 1

		for c.RainbowSkin {
			c.conn.WritePacket(&packet.PlayerSkin{
				UUID: uuid.New(),
				Skin: protocol.Skin{
					SkinID:            uuid.NewString(),
					SkinImageHeight:   32,
					SkinImageWidth:    64,
					SkinData:          getRandomSkin(),
					SkinResourcePatch: []byte("{\"geometry\": {\"default\": \"geometry.humanoid.customSlim\"}}"),
					FullID:            strconv.Itoa(i),
				},
			})
			if i == 1 {
				i = 2
			} else {
				i = 1
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
}

func getRandomSkin() []byte {
	skin := []byte{}
	randomColor := []byte{byte(rand.Intn(255)), byte(rand.Intn(255)), byte(rand.Intn(255)), 0xFF}
	for i := 0; i < 64*32; i++ {
		skin = append(skin, randomColor...)
	}

	return skin
}
