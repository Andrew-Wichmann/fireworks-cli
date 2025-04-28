#!/bin/bash

SESSION_NAME=$(basename "$PWD")

tmux new-session -dt $SESSION_NAME

tmux send-keys -t $SESSION_NAME "vim ." C-m
tmux send-keys -t $SESSION_NAME C-j

tmux new-window -t $SESSION_NAME

tmux select-window -t $SESSION_NAME:0

tmux attach-session -t $SESSION_NAME
