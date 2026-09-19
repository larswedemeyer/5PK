## Autonomous Robotics Software Stack

Software architecture for an autonomous mobile robot, developed as part of a personal robotics project and used as the topic of the German Abitur's 5. Prüfungskomponente.


## Overview

The goal is to allow a robot to perceive its environment, build an internal representation of the world and navigate autonomously.

The system is designed in a modular way so that individual components can evolve independently from simple implementations (V1) to more advanced approaches (V3).


## Architecture

The software is divided into five main modules:

# Perception
  Processes sensor data and extract relevant information about the environment.
# World Model
  Maintains an internal representation of the robot's surroundings, including detected objects, obstacles, and the goal.
# Behavior
  Determines the robot's actions based on the current state of the world model and mission objectives.
# Path Planning
  Calculates feasible paths while avoiding obstacles and optimizing movement.
# Motion Execution
  Transforms planned paths into motor commands and controls the robot's movement.


## Technologies

- Go
- Python
- C++


## Current Features

- KD-Tree
- RRT* path planning experiments
- Camera-to-Raspberry-Pi image transmission via Wi-Fi


## Planned Features

- Camera-based object and goal detection using a neural network
- Dynamic obstacle avoidance
- Autonomous navigation in new environments
- Hardware integration and testing
- Multi-sensor fusion


## Project Status

Work in progress.


## Motivation

This project serves both as a personal learning experience in robotics and artificial intelligence and as the foundation for a scientific presentation in the German Abitur.


## Author

Lars Wedemeyer
