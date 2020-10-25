import java.util.Scanner;

    public class Cipher{
        public static void main(String[] args){
            System.out.println("Digite o texto a ser criptorafado: ");
            Scanner bucky = new Scanner(System.in);
            String criptoString = bucky.nextLine();
            criptoString = criptoString.toUpperCase();
            System.out.println("Deseja criptografar com Cifra de Cesar ou com Cifra de Vigenere?");
            System.out.println("Aperte 1 para Cesar e 2 para Vigenere: ");
            int op = bucky.nextInt();
            if(op == 1){
                System.out.println("Digite a chave a ser criptorafada: ");
                int criptoKey1 = bucky.nextInt();
                String x = Cesar(criptoString, criptoKey1);
                System.out.println(x); 
            }
            if(op == 2){
                System.out.println("Digite a chave a ser criptorafada: ");
                Scanner bucky2 = new Scanner(System.in);
                String criptoKey2 = bucky2.nextLine();
                criptoKey2 = criptoKey2.toUpperCase();
                String y = Vigenere(criptoString, criptoKey2);
                System.out.println(y); 
            }
            
        
        }

        public static String Cesar(String criptoString, int key){
            char[] letters = criptoString.toCharArray();
            for(int i = 0; i < criptoString.length(); i++){
                if(criptoString.codePointAt(i) >= 65 && criptoString.codePointAt(i) <= 90){
                    int number = ((criptoString.codePointAt(i)-65+key)%26) + 65;
                    letters[i] = (char)number;
                    /*char c = char(number);
                    criptoString.charAt(i) = c;*/
                }
            }
            return new String(letters);
        }
        
        public static String Vigenere(String criptoString, String key){
            int z = 0;
            char[] letters = criptoString.toCharArray();
            //char[] letters2 = key.toCharArray();
            for(int i = 0; i < criptoString.length(); i++){
                if(criptoString.codePointAt(i) >= 65 && criptoString.codePointAt(i) <= 90){
                    int number = ((criptoString.codePointAt(i)-65+key.codePointAt(z++%key.length())-65)%26)+65;
                    letters[i] = (char)number;
                }
            }
            
            return new String(letters);
        }
        
        
    }
    