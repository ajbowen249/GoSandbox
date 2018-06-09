// IMPROVE: There is a lot of hacking in this file
//          to get it working on posix systems.
//          Most of it still holds on to Windows
//          ideas. This will need heavy refactoring
//          to make it all cleanly cross-platform.

#ifdef __WINDOWS__
#include <stdio.h>
#include <windows.h>
#include <conio.h>

const int IS_WINDOWS = 1;
void MoveTo(SHORT row, SHORT column) {
    COORD Cord;
    Cord.X = row;
    Cord.Y = column;
    SetConsoleCursorPosition(GetStdHandle(STD_OUTPUT_HANDLE), Cord);
}

char GetKey() {
    char character = 0x00;
    if(kbhit())
    {
       character = getch();
    }

    return character;
}

void SetCursorProperties(int visible) {
    CONSOLE_CURSOR_INFO cursorInfo;
    cursorInfo.dwSize = 1;
    cursorInfo.bVisible = (BOOL)visible;

    SetConsoleCursorInfo(GetStdHandle(STD_OUTPUT_HANDLE), &cursorInfo);
}

void SetCharacterProperties(int properties) {
    SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE), (WORD)properties);
}

int GetConAttributes(CONSOLE_SCREEN_BUFFER_INFO* infoBuffer) {
    return GetConsoleAttributes(GetStdHandle(STD_OUTPUT_HANDLE), infoBuffer);
}
 #else
 #include <stdio.h>
 #include <termios.h>
 #include <unistd.h>
 #include <fcntl.h>
 const int IS_WINDOWS = 0;

 // console properties
 struct termios initialTermios, currentTermios;
 int cursorDefaultVisible = 1;


void SetNoEcho() {
    currentTermios.c_lflag &= ~(ICANON | ECHO);
    tcsetattr(STDIN_FILENO, TCSANOW, &currentTermios);
}

 // Stole this from here: https://cboard.cprogramming.com/c-programming/63166-kbhit-linux.html
int kbhit() {
    int ch, oldf;

    oldf = fcntl(STDIN_FILENO, F_GETFL, 0);
    fcntl(STDIN_FILENO, F_SETFL, oldf | O_NONBLOCK);

    ch = getchar();

    fcntl(STDIN_FILENO, F_SETFL, oldf);

    if(ch != EOF) {
        ungetc(ch, stdin);
        return 1;
    }

    return 0;
 }

void MoveTo(int row, int column) {
    // terminal codes are 1-indexed.
    row++;
    column++;
    printf("\e[%i;%iH", column, row);
    fflush(stdout);
}

char GetKey() {
    char character = 0x00;
    if(kbhit()) {
       character = getchar();
       if ( character == 0x1b) {
           // Make this work like Windows
           char next = getchar(); // skip the '['
           if (next == '[') {
               return 0xE0; // return the windows special character identifier
           } else {
               ungetc(next, stdin);
           }
       }
    }

    return character;
}


void SetCursorProperties(int visible) {
    printf("\e[?25");
    printf(visible ? "h" : "l");
    fflush(stdout);
}
/*
    Hack here to translate from the Windows interface.
    Integer structure is:
    xxxxxxxxxxxxxxxxxxxxxxx|        x|           x|           x|                 xxx|                 xxx|
          reserved (blink?)|underline|bg intensity|fg intensity|background color BGR|foreground color BGR|
*/
const int FOREGROUND_RED = 0x00000001;
const int FOREGROUND_GREEN = 0x00000002;
const int FOREGROUND_BLUE = 0x00000004;
const int BACKGROUND_RED = 0x00000008;
const int BACKGROUND_GREEN = 0x00000010;
const int BACKGROUND_BLUE = 0x00000020;
const int FOREGROUND_INTENSITY = 0x00000040;
const int BACKGROUND_INTENSITY = 0x00000080;
const int COMMON_LVB_UNDERSCORE = 0x00000100;

void SetCharacterProperties(int properties) {
    int foreground = (properties & 0x00000007) + 30;
    if ( properties & FOREGROUND_INTENSITY ) {
        foreground += 60;
    }

    int background = ((properties & 0x00000038) >> 3) + 40;
    if ( properties & BACKGROUND_INTENSITY ) {
        background += 60;
    }
    printf("\e[%i;%im", foreground, background);
    fflush(stdout);
 }

void SaveInitialScreenState() {
    tcgetattr(STDIN_FILENO, &initialTermios);
    currentTermios = initialTermios;
}

void RestoreInitialScreenState() {
    tcsetattr(STDIN_FILENO, TCSANOW, &initialTermios);
    currentTermios = initialTermios;
    SetCursorProperties(cursorDefaultVisible);
}

 #endif // __WINDOWS__
