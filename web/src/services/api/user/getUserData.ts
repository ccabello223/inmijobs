import type { UserModel } from '@/models/UserModel';
import { usersMock } from '../../mock/usersMock';

export const getUserData = async (id: string): Promise<UserModel | null> => {
	if (id) {
		return usersMock.find(u => u.ID === id) ?? null;
	}
	return usersMock[0] ?? null;
};