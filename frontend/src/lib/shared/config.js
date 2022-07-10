import ky from 'ky';

export async function config() {
	return import.meta.env.DEV ? dev() : await prod();
}

function dev() {
	return {
		apiUrl: import.meta.env.VITE_API_URL,
		notificationsUrl: import.meta.env.VITE_NOTIFICATIONS_URL,
		auth: {
			url: import.meta.env.VITE_AUTH_URL,
			realm: import.meta.env.VITE_AUTH_REALM,
			clientId: import.meta.env.VITE_AUTH_CLIENT_ID
		},
		profileUrl: import.meta.env.VITE_PROFILE_URL,
		dashboardUrl: import.meta.env.VITE_DASHBOARD_URL
	};
}

function prod() {
	return ky.get('/config').json();
}
