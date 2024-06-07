import { InputAdornment, Typography } from "@mui/material";
import TextField, { TextFieldProps } from "@mui/material/TextField";
import React from "react";

type IconTextFieldProps = TextFieldProps & {
  iconStart?: React.ReactNode;
  iconEnd?: React.ReactNode;
};

interface HeadingTextProps {
  text: string;
}

export const IconTextField: React.FC<IconTextFieldProps> = ({
  iconStart,
  iconEnd,
  ...props
}) => {
  return (
    <TextField
      {...props}
      sx={{ width: "100%" }}
      InputProps={{
        startAdornment: iconStart ? (
          <InputAdornment position="start">{iconStart}</InputAdornment>
        ) : null,
        endAdornment: iconEnd ? (
          <InputAdornment position="end">{iconEnd}</InputAdornment>
        ) : null,
      }}
    />
  );
};

export const HeadingText: React.FC<HeadingTextProps> = ({ text }) => {
  return (
    <Typography
      variant="h2"
      gutterBottom
      sx={{
        width: "fit-content",
        margin: " 1rem auto",
        fontWeight: "bold",
        color: "#6358DC",
      }}
    >
      {text}
    </Typography>
  );
};
