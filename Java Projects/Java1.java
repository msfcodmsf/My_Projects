import java.util.Scanner;

public class Java1 {
    public static void main(String[] args) {
        System.out.println("Lütfen 1-10 arasında bir sayı giriniz.");
        Random random = new Random(); // Tr: Random sınıfından bir nesne oluşturur. // En: Creates an object of class Random.
        int guessNumber = random.nextInt(10) + 1; // Tr: 1 ile 10 arasında random sayı oluşturma  // En: Generate a random number between 1 and 10
        Scanner scanner = new Scanner(System.in); // Tr: Scanner sınıfından bir nesne oluşturur. // En: Creates an object of class Scanner.
        System.out.print(">>");
        int number = scanner.nextInt(); // Tr: Kullanıcıdan veri alma işlemi // En: Retrieving data from the user
        // Tr: 1. Tahmin işlemi // En: 1. Estimation process
            if (number > guessNumber) {
                System.out.println("Lütfen daha küçük bir sayı giriniz. Kalan tahmin hakkınız:3");
                System.out.print(">>");
                number = scanner.nextInt();
            } else if (number < guessNumber) {
                System.out.println("Lütfen daha büyük bir sayı giriniz. Kalan tahmin hakkınız:3");
                System.out.print(">>");
                number = scanner.nextInt();
            }else if (number == guessNumber){
                System.out.println("Sayıyı buldunuz. Kazandınız.");
            }
        // Tr: 2. Tahmin işlemi // En: 2. Estimation process
            if (number > guessNumber) {
                System.out.println("Lütfen daha küçük bir sayı giriniz. Kalan tahmin hakkınız:2");
                System.out.print(">>");
                number = scanner.nextInt();
            } else if (number < guessNumber) {
                System.out.println("Lütfen daha büyük bir sayı giriniz. Kalan tahmin hakkınız:2");
                System.out.print(">>");
                number = scanner.nextInt();
            }else if (number == guessNumber){
                System.out.println("Sayıyı buldunuz. Kazandınız.");
            }
        // Tr: 3. Tahmin işlemi // En: 3. Estimation process
            if (number > guessNumber) {
                System.out.println("Lütfen daha küçük bir sayı giriniz. Kalan tahmin hakkınız:1");
                System.out.print(">>");
                number = scanner.nextInt();
            } else if (number < guessNumber) {
                System.out.println("Lütfen daha büyük bir sayı giriniz. Kalan tahmin hakkınız:1");
                System.out.print(">>");
                number = scanner.nextInt();
            }else if (number == guessNumber){
                System.out.println("Sayıyı buldunuz. Kazandınız.");
            }

        // Tr: SON TAHMİN HAKKI OLDUĞU İÇİN DAHA FAZLA İŞLEME İZİN VERİLMİYOR
        // En: NO FURTHER PROCESSING IS ALLOWED AS IT IS THE LAST GUESS
            if (number == guessNumber) {
                System.out.println("Sayıyı buldunuz. Kazandınız.");
            } else {
                System.out.println("Kaybettiniz Bilemediğiniz sayi:"+ guessNumber);
                // (Bilemediginiz sayi:"+ guessNumber) Üsame abiden çaldım :)
            }
    }
}