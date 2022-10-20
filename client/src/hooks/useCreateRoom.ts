import axios, { AxiosError, AxiosResponse } from 'axios';
import { useState } from 'react';
import { Room } from '../@types';

interface UseCreateRoomResponse {
  isLoading: boolean;
  createNewRoom: () => Promise<Room & { err: AxiosError | null }>;
}

const promHandler = async <T>(
  prom: Promise<T>
): Promise<[Awaited<T> | null, AxiosError | null]> => {
  try {
    return [await prom, null];
  } catch (error) {
    return [null, error as AxiosError];
  }
};

export const useCreateRoom = (): UseCreateRoomResponse => {
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const createNewRoom = async () => {
    setIsLoading(true);

    const [newRoom, err] = await promHandler<AxiosResponse<Room>>(
      axios.post(
        `http://${import.meta.env.VITE_API_ENDPOINT}/api/create`,
        {},
        { headers: { 'Access-Control-Allow-Headers': '*' } }
      )
    );

    setIsLoading(false);

    return {
      roomId: (newRoom as AxiosResponse<Room>)?.data?.roomId ?? null,
      err
    };
  };

  return { isLoading, createNewRoom };
};
