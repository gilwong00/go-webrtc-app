package room

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/gilwong00/go-webrtc-app/pkg/models"
	"github.com/gorilla/websocket"
)

type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]models.Participant
}

// initialize RoomMap
func (r *RoomMap) Init() {
	r.Map = make(map[string][]models.Participant)
}

// gets all participants in room
func (r *RoomMap) GetRoom(roomId string) []models.Participant {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	return r.Map[roomId]
}

// creates new chat room
func (r *RoomMap) CreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	// creates random room id
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghojklmnopqrstuvwxyz123456789ABCDEFGHIJKLMNPQRSTUVWXYZ")
	b := make([]rune, 8)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	roomId := string(b)
	r.Map[roomId] = []models.Participant{}

	return roomId
}

// Inserts participant into room
func (r *RoomMap) AddParticipantToRoom(roomId string, isHost bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	p := models.Participant{
		Host: isHost,
		Conn: conn,
	}

	log.Println("Add new participant into room with id: ", roomId)
	r.Map[roomId] = append(r.Map[roomId], p)
}

// delete room
func (r *RoomMap) DeleteRoom(roomId string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	delete(r.Map, roomId)
}
