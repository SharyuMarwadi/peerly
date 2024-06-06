import { Box, Divider, Typography } from "@mui/material";
import loginImage from '../../assets/LoginImage.png'
import LoginForm from "./LoginForm";


const LoginPage = () => {
    return (
        <Box
            sx={{
                display: 'flex',
                flexDirection: 'row',
                alignContent: 'space-between',
                margin: 'auto'
            }}
        >
            <Box>
                <Box component="img" sx={{
                    width: 650,
                    height: 500,
                    // maxHeight: { xs: 233, md: 167 },
                    // maxWidth: { xs: 350, md: 250 },
                }}
                    alt="Login image"
                    src={loginImage}
                />
                <Typography variant="h1" gutterBottom sx={{ width: 'fit-content', margin: 'auto' }}>
                    Peerly
                </Typography>
                <Typography variant="subtitle1" gutterBottom sx={{ width: 'fit-content', margin: 'auto' }}>
                    The peer rewards and recognition system
                </Typography>

            </Box>
            <Divider orientation="vertical" variant="middle" flexItem />
            <Box sx={{ width: '30%' }}>
                <LoginForm />
            </Box>

        </Box>
    );
}
export default LoginPage;



