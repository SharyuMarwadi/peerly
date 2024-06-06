import { InputAdornment } from "@mui/material";
import TextField, { TextFieldProps } from "@mui/material/TextField";
import React from "react";


type IconTextFieldProps = TextFieldProps & {
    iconStart?: React.ReactNode;
    iconEnd?: React.ReactNode;
}

export const IconTextField: React.FC<IconTextFieldProps> = ({ iconStart, iconEnd, ...props }) => {
    return (
        <TextField
            {...props}
            sx={{width: 'maxWidth'}}
            InputProps={{
                startAdornment: iconStart ? (
                    <InputAdornment position="start">{iconStart}</InputAdornment>
                ) : null,
                endAdornment: iconEnd ? (
                    <InputAdornment position="end">{iconEnd}</InputAdornment>
                ) : null
            }}
        />
    );
};
