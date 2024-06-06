import Box from '@mui/material/Box';
import EmailIcon from '@mui/icons-material/Email';
import KeyIcon from '@mui/icons-material/Key';
import { IconTextField } from '../../shared/wrapper';
import Divider from '@mui/material/Divider';
import GoogleLoginBox from './GoogleLoginBox';
import { Button } from '@mui/material';

const LoginForm = () => {
    return (
        <>
            <Box
                component="form"
                sx={{
                    '& > :not(style)': { m: 1, width: 'auto' },
                    display: 'flex',
                    flexDirection: 'column'
                }}
                noValidate
                autoComplete="off"
            >
                <Divider>OR</Divider>
                <IconTextField
                    label="Email"
                    iconStart={<EmailIcon />}
                />
                <IconTextField
                    label="Password"
                    iconStart={<KeyIcon />}
                />
                <Button variant="contained" sx={{ width: '25ch' }}>LogIn</Button>
            </Box>
        </>
    )
}
export default LoginForm


                // <GoogleLoginBox />
