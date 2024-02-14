@echo off

IF NOT EXIST bin mkdir bin

go build -o bin\pico.exe