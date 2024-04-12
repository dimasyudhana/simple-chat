import { useState, useEffect, useContext } from 'react';
import { AuthContext } from './context/AuthProvider';
import { WebsocketContext } from './context/WebsocketProvider';
import { v4 as uuidv4 } from 'uuid';
import { useNavigate } from 'react-router-dom';

const Homepage = () => {
  const [rooms, setRooms] = useState([]);
  const [roomName, setRoomName] = useState('');
  const { user } = useContext(AuthContext);
  const { setConn } = useContext(WebsocketContext);
  const navigate = useNavigate();

  const getRooms = async () => {
    try {
      const res = await fetch(`${'http://127.0.0.1:8080'}/rooms`, {
        method: 'GET',
      });

      if (res.ok) {
        const data = await res.json();
        setRooms(data);
      }
    } catch (err) {
      console.log(err);
    }
  };

  useEffect(() => {
    getRooms();
  }, []);

  const submitHandler = async (e) => {
    e.preventDefault();

    try {
      setRoomName('');
      const res = await fetch(`${'http://127.0.0.1:8080'}/rooms/register`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify({
          id: uuidv4(),
          name: roomName,
        }),
      });

      if (res.ok) {
        getRooms();
        navigate('/homepage');
      }      
    } catch (err) {
      console.log(err);
    }
  };

  const joinRoom = (roomId) => {
    const ws = new WebSocket(
      `${'ws://127.0.0.1:8080'}/rooms/:room_id/join${roomId}?userId=${user.id}&username=${user.username}`
    );
    if (ws.readyState === ws.OPEN) {
      setConn(ws);
    }
  };

  return (
    <div className='my-8 px-4 md:mx-32 w-full h-full'>
      <div className='flex justify-center mt-3 p-5'>
        <input
          type='text'
          className='border border-grey p-2 rounded-md focus:outline-none focus:border-blue'
          placeholder='room name'
          value={roomName}
          onChange={(e) => setRoomName(e.target.value)}
        />
        <button
          className='bg-blue border text-white rounded-md p-2 md:ml-4'
          onClick={submitHandler}
        >
          create room
        </button>
      </div>
      <div className='mt-6'>
        <div className='font-bold'>Available Rooms</div>
        <div className='grid grid-cols-1 md:grid-cols-5 gap-4 mt-6'>
          {rooms.map((room) => (
            <div
              key={room.id}
              className='border border-blue p-4 flex items-center rounded-md w-full'
            >
              <div className='w-full'>
                <div className='text-sm'>room</div>
                <div className='text-blue font-bold text-lg'>{room.name}</div>
              </div>
              <div className=''>
                <button
                  className='px-4 text-white bg-blue rounded-md'
                  onClick={() => joinRoom(room.id)}
                >
                  join
                </button>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default Homepage;
