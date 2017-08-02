#!/usr/bin/env bash
docker rmi ivanthescientist/tournament_service
docker build -t ivanthescientist/tournament_service .
docker push ivanthescientist/tournament_service