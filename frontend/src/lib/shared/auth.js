import Keycloak from 'keycloak-js';
import { stores } from './stores.js';

// About the 'keycloak-js' eval warning during build, see:
// https://github.com/keycloak/keycloak/issues/10710

export default async function auth() {
	/** @type {Keycloak} */
	let keycloak;

	function isSignedIn() {
		return keycloak.authenticated;
	}

	async function signIn() {
		try {
			if (!isSignedIn()) {
				await keycloak.login();
			}
		} catch (err) {
			console.error(err);
			throw new Error('error signing in');
		}
	}

	async function signOut() {
		try {
			if (isSignedIn()) {
				await keycloak.logout({
					redirectUri: location.origin
				});
			}
		} catch (err) {
			console.error(err);
			throw new Error('error signing out');
		}
	}

	function getUserName() {
		return keycloak.userInfo?.name;
	}

	async function getAccessToken() {
		try {
			await keycloak.updateToken(null);
			return keycloak.token;
		} catch (err) {
			console.error(err);
			throw new Error('error getting access token');
		}
	}

	async function init() {
		try {
			keycloak = new Keycloak(stores.config.auth);

			await keycloak.init({
				pkceMethod: 'S256',
				onLoad: 'login-required',
				checkLoginIframe: false
			});

			if (keycloak.authenticated) {
				await keycloak.loadUserInfo();
			}
		} catch (err) {
			console.error(err);
			throw new Error('error initializing auth');
		}
	}

	await init();

	return {
		isSignedIn,
		signIn,
		signOut,
		getUserName,
		getAccessToken
	};
}
