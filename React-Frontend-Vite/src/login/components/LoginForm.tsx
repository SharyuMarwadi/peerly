import Box from "@mui/material/Box";
import EmailIcon from "@mui/icons-material/Email";
import KeyIcon from "@mui/icons-material/Key";
import { HeadingText, IconTextField } from "../../shared/wrapper";
import Divider from "@mui/material/Divider";
import GoogleLoginBox from "./GoogleLoginBox";
import { Button, Typography } from "@mui/material";

const LoginForm = () => {
  return (
    <>
      <Box
        component="form"
        sx={{
          "& > :not(style)": { m: 1 },
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
        noValidate
        autoComplete="off"
      >
        <Box
          sx={{
            display: "block",
            width: "25rem",
          }}
        >
          <HeadingText text="Welcome"/>
        </Box>
        <Box sx={{ width: "100%" }}>
          <GoogleLoginBox />
        </Box>
        <Divider sx={{ width: "100%" }}>OR</Divider>
        <IconTextField
          label="Email"
          iconStart={<EmailIcon />}
          sx={{ width: "100%" }}
        />
        <IconTextField
          label="Password"
          iconStart={<KeyIcon />}
          sx={{ width: "100%" }}
        />
        <Box
          sx={{
            width: "100%",
            display: "flex",
            justifyContent: "end",
            margin: "0",
            marginTop: "1rem",
          }}
        >
          <Typography variant="caption" display="block" gutterBottom>
            Forgot Password?
          </Typography>
        </Box>
        <Button variant="contained" sx={{ width: "100%", backgroundColor: '#6358DC' }}>
          LogIn
        </Button>
      </Box>
    </>
  );
};
export default LoginForm;

