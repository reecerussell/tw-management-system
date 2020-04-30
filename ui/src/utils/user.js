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

const listeners = new Map();

const Listen = (name, callback) => listeners.set(name, callback);
const Unlisten = (name) => listeners.delete(name);

const triggerListeners = () => listeners.forEach((c) => c());

const GetId = () => getCurrentPayload()["user_id"];

const getCurrentPayload = () => {
	const token = GetAccessToken();
	if (!token) {
		return null;
	}

	const parts = token.split(".");
	if (parts.length < 2) {
		return null;
	}

	const payloadData = atob(parts[1]);

	return JSON.parse(payloadData);
};

export {
	GetAccessToken,
	Login,
	Logout,
	IsAuthenticated,
	Listen,
	Unlisten,
	GetId,
};
