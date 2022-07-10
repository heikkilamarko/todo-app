import Keycloak from 'keycloak-js';
import { stores } from './stores.js';

let keycloak;

export async function init() {
	try {
		keycloak = new Keycloak(stores.config.auth);
		const isAuthenticated = await keycloak.init({
			pkceMethod: 'S256'
			// enableLogging: true
		});
		if (isAuthenticated) {
			await keycloak.loadUserInfo();
		} else {
			await keycloak.login();
		}
		return isAuthenticated;
	} catch (err) {
		console.log(err);
		throw new Error('auth init failed');
	}
}

export async function accessToken() {
	try {
		await keycloak.updateToken(null);
		return keycloak.token;
	} catch (error) {
		await keycloak.login();
	}
}

export async function logout() {
	await keycloak.logout();
}

export function userName() {
	return keycloak.userInfo?.name ?? '<unknown user>';
}

export function isUserRole() {
	return isInRole('todo-user');
}

export function isViewerRole() {
	return isInRole('todo-viewer');
}

export function isInRole(role) {
	return keycloak.hasResourceRole(role, 'todo-api');
}
