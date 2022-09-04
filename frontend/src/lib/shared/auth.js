import Keycloak from 'keycloak-js';
import { stores } from './stores.js';

/** @type {Keycloak} */
let keycloak;

export const Roles = {
	User: 'todo-user',
	Viewer: 'todo-viewer'
};

export async function initAuth() {
	try {
		keycloak = new Keycloak(stores.config.auth);

		await keycloak.init({
			pkceMethod: 'S256',
			onLoad: 'login-required'
		});

		if (keycloak.authenticated) {
			await keycloak.loadUserInfo();
		}
	} catch (err) {
		console.error(err);
		throw new Error('error initializing auth');
	}
}

export function isSignedIn() {
	return keycloak.authenticated;
}

export async function signIn() {
	try {
		if (!isSignedIn()) {
			await keycloak.login();
		}
	} catch (err) {
		console.error(err);
		throw new Error('error signing in');
	}
}

export async function signOut() {
	try {
		if (isSignedIn()) {
			await keycloak.logout();
		}
	} catch (err) {
		console.error(err);
		throw new Error('error signing out');
	}
}

export async function getAccessToken() {
	try {
		await keycloak.updateToken(null);
		return keycloak.token;
	} catch (err) {
		console.error(err);
		throw new Error('error getting access token');
	}
}

export function getUserName() {
	return keycloak.userInfo?.name;
}

export function isInRole(role, resource = 'todo-api') {
	return keycloak.hasResourceRole(role, resource);
}
