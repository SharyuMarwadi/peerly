import { GoogleLogin } from '@react-oauth/google';
import { jwtDecode } from "jwt-decode";

const GoogleLoginBox = () => {
    return (
        <GoogleLogin
            onSuccess={credentialResponse => {
                if (credentialResponse.credential) {
                    const credentialResponseDecoded = jwtDecode(credentialResponse.credential);
                    console.log(credentialResponseDecoded);
                } else {
                    console.log("Login Failed");
                }
            }}
            onError={() => {
                console.log('Login Failed');
            }}
        />
    );
}

export default GoogleLoginBox;
