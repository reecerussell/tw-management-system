const AccessTokenCookieName = "twms_ac";

const GetAccessToken = () => {
	const value = document.cookie.match(
		"(^|;) ?" + AccessTokenCookieName + "=([^;]*)(;|$)"
	);

	return value ? value[2] : null;
};

const Login = (token, expires) => {
	expires *= 1000;
	let d = new Date(expires);
	d = d + d.getTimezoneOffset() * 60000;

	document.cookie =
		AccessTokenCookieName + "=" + token + ";path=/;expires=" + d;

	triggerListeners();
};

const Logout = () => {
	Login(null, -1);

	triggerListeners();
};

const IsAuthenticated = () => {
	return GetAccessToken() !== null;
};

const listeners = [];

const Listen = (callback) => listeners.push(callback);

const triggerListeners = () => listeners.forEach((c) => c());

export { GetAccessToken, Login, Logout, IsAuthenticated, Listen };
