import { dev } from '$app/environment';
import * as env from '$env/static/public';
import ky from 'ky';

export default async function config() {
	return dev ? devConfig() : await prodConfig();
}

function devConfig() {
	return {
		apiUrl: env.PUBLIC_API_URL,
		notificationsUrl: env.PUBLIC_NOTIFICATIONS_URL,
		auth: {
			url: env.PUBLIC_AUTH_URL,
			realm: env.PUBLIC_AUTH_REALM,
			clientId: env.PUBLIC_AUTH_CLIENT_ID
		},
		profileUrl: env.PUBLIC_PROFILE_URL,
		dashboardUrl: env.PUBLIC_DASHBOARD_URL
	};
}

function prodConfig() {
	return ky.get('/config').json();
}
