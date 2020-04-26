const AccessTokenCookieName = "mpp_ac";

const GetAccessToken = () => {
	const value = document.cookie.match(
		"(^|;) ?" + AccessTokenCookieName + "=([^;]*)(;|$)"
	);

	return value ? value[2] : null;
};

const Login = (token, expires) => {
	const d = new Date(expires);
	document.cookie =
		AccessTokenCookieName +
		"=" +
		token +
		";path=/;expires=" +
		d.toGMTString();
};

const Logout = () => {
	SetAccessToken(null, -1);
};

const IsAuthenticated = () => {
	return GetAccessToken() !== null;
};

export { GetAccessToken, Login, Logout, IsAuthenticated };
