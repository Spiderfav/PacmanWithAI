# Final Year Project

This repository has been created to store your final year project.

You may edit it as you like, but please do not remove the default topics or the project members list. These need to stay as currently defined in order for your supervisor to be able to find your project.

For the markers of the Project:

    - I have realised that some of my commits may have an account name as "Spiderfav" or "Rui Favinha". 
    
Now I don't exactly know how since my machines are using the access tokens created to commit the files. I've tried to change it now but if any commit appears as "Spiderfav" or "Rui Favinha", it really should appear as "Favinha Rui (2021) ZKAC432". "Spiderfav" is just my normal GitHub and online user handle.

-------------------

# This is a simple guide to getting this program to run on your machine.

1. Make sure you have the latest Golang Version Installed

2. On a terminal of your chosen Operating System, navigate the terminal to be in the folder where "main.go" file is found

3. Run the following command on the terminal window:
    - go build -o Maze

    This will compile the code into a binary file to be easily run.
    It should take only a second or two to fully compile

4. On the same terminal, run the command:
    - ./Maze (Linux/MacOS)
    - run Maze (Windows)

    The program should begin to run in the a window or fullscreen (depending on your monitor's resolution)

5. To use the program, the following keys are used:
    1 - Switch to smallest Maze Size
    2 - Switch to medium Maze Size
    3 - Switch to largest Maze Size
    A - Run's Dijkstras Algorithm on the Maze
    B - Run's A* Algorithm on the Maze
    C - Run's "Maze to Weighted Graph" Algorithm on the Maze
    D - Clear's the screen to the normal maze built

    ### NOTE: Clicking on 1-3 multiple times will result in new maze configurations each time!

6. While the maze is being run, the terminal window used to run the maze will also output the time taken for each algorithm to run on the disired maze size.

7. To quit the program, simply "Alt+Tab" to the terminal window used to run the program and "Ctrl+C" the terminal window

