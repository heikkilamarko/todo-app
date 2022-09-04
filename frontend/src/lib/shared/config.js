import ky from 'ky';

export async function config() {
	return import.meta.env.DEV ? dev() : await prod();
}

function dev() {
	return {
		apiUrl: import.meta.env.PUBLIC_API_URL,
		notificationsUrl: import.meta.env.PUBLIC_NOTIFICATIONS_URL,
		auth: {
			url: import.meta.env.PUBLIC_AUTH_URL,
			realm: import.meta.env.PUBLIC_AUTH_REALM,
			clientId: import.meta.env.PUBLIC_AUTH_CLIENT_ID
		},
		profileUrl: import.meta.env.PUBLIC_PROFILE_URL,
		dashboardUrl: import.meta.env.PUBLIC_DASHBOARD_URL
	};
}

function prod() {
	return ky.get('/config').json();
}
