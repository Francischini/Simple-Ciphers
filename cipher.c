#include <stdio.h>
#include <stdlib.h>

void cesar(char *m, int n, int k) {
  int i;
  for (i = 0; i < n; ++i)
    if (m[i] >= 'A' && m[i] <= 'Z')
      m[i] = (m[i]-'A'+k)%26+'A';
}

void vigenere(char *m, int n, char *c, int s) {
  int i, z = 0;
  for (i = 0; i < n; ++i)
    if (m[i] >= 'A' && m[i] <= 'Z')
      m[i] = (m[i]-'A'+c[z++%s]-'A')%26+'A';
}

int main(){
    int n, op, s, k;
    char *m, *c;

    scanf("%d", &n);
    m = malloc(n+1);

    scanf(" %[^\n]%d", m, &op);

    if (op == 1) {
      scanf("%d", &k);
      cesar(m, n, k);
    }
    else {
      scanf("%d", &s);
      c = malloc(s+1);
      scanf(" %[^\n]", c);
      vigenere(m, n, c, s);
    }

    printf("%s\n", m);

    return 0;
}

//Special thanks to Sergio
