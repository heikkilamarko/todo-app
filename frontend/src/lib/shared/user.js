import { stores } from './stores.js';
import * as api from './api.js';

export default async function user() {
	const { auth } = stores;

	const { data: userinfo } = await api.getUserinfo();

	function username() {
		return auth.getUserName();
	}

	function hasPermission(p) {
		return userinfo?.permissions?.includes(p) ?? false;
	}

	return { username, hasPermission };
}
