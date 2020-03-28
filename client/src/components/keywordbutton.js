import React from 'react';
import Button from '@material-ui/core/Button';
import { makeStyles } from '@material-ui/core/styles';
import DeleteIcon from '@material-ui/icons/Delete';
import IconButton from '@material-ui/core/IconButton';
import Typography from '@material-ui/core/Typography';

const useStyles = makeStyles(theme => ({
    root: {
      backgroundColor: 'white',
    },
  }));
  
  
  export default function RedditKeywords(props) {
    const classes = useStyles();

    var keywordButtons = []

    for (var i = 0; i < props.values.length; i++) {
        keywordButtons.push(
            <Button className={classes.root}>
                <Typography>{props.values[i]}</Typography>
                <IconButton aria-label="delete">
                    <DeleteIcon />
                </IconButton>

                <IconButton aria-label="delete">
                    <DeleteIcon />
                </IconButton>
            </Button>
        )
      }
  
    return (
        keywordButtons
    );
  }
