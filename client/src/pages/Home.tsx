import { AxiosError } from 'axios';
import { useCallback, useState } from 'react';
import { useNavigate } from 'react-router-dom';

import { useCreateRoom } from '../hooks';

export const Home = () => {
  const [createError, setCreateError] = useState<AxiosError | null>(null);
  const { isLoading, createNewRoom } = useCreateRoom();
  const navigate = useNavigate();

  const handleCreateRoom = useCallback(async () => {
    const { roomId, err } = await createNewRoom();

    if (err) return setCreateError(err as AxiosError);
    return navigate(`/room/${roomId}`);
  }, []);

  if (isLoading) return <div>Loading.....</div>;
  if (createError) return <div>{createError.message}</div>;

  return (
    <div>
      <button onClick={handleCreateRoom}>Create Room</button>
    </div>
  );
};
