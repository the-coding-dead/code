#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include <dirent.h>
#include <fcntl.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>

#define DIRSIZ 50

char *fmtname(char *path) {
  char *p;

  // Find first character after last slash.
  for (p = path + strlen(path); p >= path && *p != '/'; p--)
    ;
  p++;

  // Return blank-padded name.
  if (strlen(p) >= DIRSIZ)
    return p;

  static char buf[DIRSIZ + 1];
  memmove(buf, p, strlen(p));
  memset(buf + strlen(p), ' ', DIRSIZ - strlen(p));
  return buf;
}

void ls(char *path) {
  struct stat st;

  if (stat(path, &st) < 0) {
    printf("ls: cannot stat %s\n", path);
    return;
  }

  if (S_ISREG(st.st_mode)) {
    printf("%s %d %lu %ld\n", fmtname(path), st.st_mode, st.st_ino, st.st_size);
  }

  if (S_ISDIR(st.st_mode)) {
    char buf[512];
    if (strlen(path) + 1 + DIRSIZ + 1 > sizeof buf) {
      printf("ls: path too long\n");
      return;
    }
    strcpy(buf, path);
    char *p = buf + strlen(buf);
    *p++ = '/';

    DIR *dir = opendir(path);
    for (struct dirent *de = readdir(dir); de != NULL; de = readdir(dir)) {
      memmove(p, de->d_name, DIRSIZ);
      p[DIRSIZ] = 0;
      if (stat(buf, &st) < 0) {
        printf("ls: cannot stat %s\n", buf);
        continue;
      }
      printf("%s %d %lu %ld\n", fmtname(buf), st.st_mode, st.st_ino,
             st.st_size);
    }
  }
}

int main(int argc, char *argv[]) {
  if (argc < 2) {
    ls(".");
    exit(0);
  }

  for (int i = 1; i < argc; i++)
    ls(argv[i]);

  exit(0);
}
