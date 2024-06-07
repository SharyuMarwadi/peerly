import { Box, Divider, Typography } from "@mui/material";
import loginImage from '../../assets/LoginImage.png'
import LoginForm from "./LoginForm";
import { HeadingText } from "../../shared/wrapper";


const LoginPage = () => {
    return (
        <Box
            sx={{
                display: 'flex',
                flexDirection: 'row',
                alignContent: 'space-between',
                justifyContent: 'space-evenly',
                margin: '6rem auto'
            }}
        >
            <Box>
                <Box component="img" sx={{
                    width: 550,
                    height: 400,
                }}
                    alt="Login image"
                    src={loginImage}
                />
                <HeadingText text="Peerly"/>
                <Typography variant="subtitle1" gutterBottom sx={{ width: 'fit-content', margin: 'auto', fontWeight: 'bold' }}>
                    The peer rewards and recognition system
                </Typography>

            </Box>
            <Divider orientation="vertical" variant="middle" flexItem />
            <Box sx={{ 
                width: '30%',
                display: 'flex',
                alignItems: 'center' 
                }}>
                <LoginForm />
            </Box>

        </Box>
    );
}
export default LoginPage;



