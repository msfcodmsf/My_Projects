import java.util.Scanner;

public class Restaurant_Menu {
    public static void main(String[] args) {
        // Çorba
        System.out.println("✨Restoranımıza Hoşgeldiniz✨");
        System.out.println("Çorba Kategorilerimiz.");
        System.out.println("1- Mercimek Çorbası -Fiyat: 10 TL");
        System.out.println("2- Domates Çorbası -Fiyat: 12 TL");
        System.out.println("3- Tavuk Çorbası -Fiyat: 15 TL");
        System.out.println("4- Çorba İstemiyorum.");
        System.out.println("Numara girmeniz yeterlidir.");
        System.out.print("->");
        Scanner scanner = new Scanner(System.in); // Tr: Scanner sınıfından bir nesne oluşturur. // En: Creates an object of class Scanner.
        int corbalar = scanner.nextInt(); // Tr: Kullanıcıdan veri alma işlemi // En: Retrieving data from the user.
        if (corbalar >= 5 || corbalar <= 10000) {
            System.out.println("Yanlış sayı girdiniz tekrardan deneyiniz");
            System.out.println("Çorba Kategorilerimiz.");
            System.out.println("1- Mercimek Çorbası -Fiyat: 10 TL");
            System.out.println("2- Domates Çorbası -Fiyat: 12 TL");
            System.out.println("3- Tavuk Çorbası -Fiyat: 15 TL");
            System.out.println("4- Çorba İstemiyorum.");
            System.out.println("Numara girmeniz yeterlidir.");
            System.out.print("->");
            corbalar = scanner.nextInt();
        }

        int anaYemekler = 0;
        if (corbalar > 0 && corbalar < 5) {
            // Ana Yemek bölümü
            System.out.println("1- Izgara Tavuk - 25 TL");
            System.out.println("2- Köfte - 20 TL");
            System.out.println("3- Sebzeli Makarna - 18 TL");
            System.out.println("4- Ana Yemek İstemiyorum.");
            System.out.println("Numara girmeniz yeterlidir.");
            System.out.print("->");
            anaYemekler = scanner.nextInt();
        } if (anaYemekler > 5 || anaYemekler < 10000) {
            System.out.println("Yanlış sayı girdiniz tekrardan deneyiniz");
            // Ana Yemek bölümü
            System.out.println("1- Izgara Tavuk - 25 TL");
            System.out.println("2- Köfte - 20 TL");
            System.out.println("3- Sebzeli Makarna - 18 TL");
            System.out.println("4- Ana Yemek İstemiyorum.");
            System.out.println("Numara girmeniz yeterlidir.");
            System.out.print("->");
            anaYemekler = scanner.nextInt();
        }
        int tatli = 0;
        if (anaYemekler > 0 && anaYemekler < 5) {
            // Tatlı bölümü
            System.out.println("1- İrmik Helvası -Fiyat: 8 TL");
            System.out.println("2- Baklava -Fiyat: 12 TL");
            System.out.println("3- Dondurma -Fiyat: 10 TL");
            System.out.println("4- Tatlı yemek İstemiyorum.");
            System.out.println("Numara girmeniz yeterlidir.");
            System.out.print("->");
            tatli = scanner.nextInt();
        } if (tatli >= 5 || tatli <= 10000) {
            System.out.println("Yanlış sayı girdiniz tekrardan deneyiniz");
            // Tatlı bölümü
            System.out.println("1- İrmik Helvası -Fiyat: 8 TL");
            System.out.println("2- Baklava -Fiyat: 12 TL");
            System.out.println("3- Dondurma -Fiyat: 10 TL");
            System.out.println("4- Tatlı yemek İstemiyorum.");
            System.out.println("Numara girmeniz yeterlidir.");
            System.out.print("->");
            tatli = scanner.nextInt();
        }


        // Çorba fiyatları
        int mercimek;
        int domates;
        int tavuk;
        // Ana yemekler fiyatları
        int sebze;
        int izgara;
        int kofte;
        // Tatlı fiyatları
        int irmik;
        int baklava;
        int dondurma;

        if ((corbalar > 0 && corbalar < 5) && (anaYemekler > 0 && anaYemekler < 5) && ((tatli > 0 && tatli < 5))) {
            System.out.println("Hesap Gösterimi");
            int corbalarToplam = 0;
            switch (corbalar) {
                case 1:
                    mercimek = 10;
                    System.out.println("Mercimek Çorbası " + mercimek + "Tl");
                    corbalarToplam = mercimek;
                    break;
                case 2:
                    domates = 12;
                    System.out.println("Domates Çorbası " + domates + "Tl");
                    corbalarToplam = domates;
                    break;
                case 3:
                    tavuk = 15;
                    System.out.println("Tavuk Çorbası " + tavuk + "Tl");
                    corbalarToplam = tavuk;
                    break;
            }
            int anaYemeklerToplam = 0;
            switch (anaYemekler) {
                case 1:
                    izgara = 25;
                    System.out.println("Izgara Tavuk Ana Yemek " + izgara + "Tl");
                    anaYemeklerToplam = izgara;
                    break;
                case 2:
                    kofte = 20;
                    System.out.println("Köfte Ana Yemek " + kofte + "Tl");
                    anaYemeklerToplam = kofte;
                    break;
                case 3:
                    sebze = 18;
                    System.out.println("Sebzeli Makarna Ana Yemek " + sebze + "Tl");
                    anaYemeklerToplam = sebze;
                    break;
            }
            int tatlitoplam = 0;
            switch (tatli) {
                case 1:
                    irmik = 8;
                    System.out.println("İrmik Helvası Tatlısı " + irmik + "Tl");
                    tatlitoplam = irmik;
                    break;
                case 2:
                    baklava = 12;
                    System.out.println("Baklava Tatlısı " + baklava + "Tl");
                    tatlitoplam = baklava;
                    break;
                case 3:
                    dondurma = 10;
                    System.out.println("Dondurma Tatlısı " + dondurma + "Tl");
                    tatlitoplam = dondurma;
                    break;
            }
            // %20 kdv oranı
            double kdv = 0.20;
            double kdvOrani;
            // %10 kvr oranı
            double kvr = 0.10;
            double kvrOrani;

            int anaToplam = corbalarToplam + anaYemeklerToplam + tatlitoplam;
            kvrOrani = anaToplam * kvr;
            kdvOrani = anaToplam * kdv;
            double sonToplam = kvrOrani + kdvOrani + anaToplam;
            System.out.println("Küver Ücreti (%10): " + kvrOrani + "Tl");
            System.out.println("KDV (%20): " + kdvOrani + "Tl");
            System.out.println("Toplam Ödeme: " + sonToplam + "Tl");
            System.out.println("✨Tekrar Görüşmek Üzere✨");
        } else if ((corbalar > 5 || corbalar < 10000) && (anaYemekler > 5 || anaYemekler < 10000) && (tatli > 5 || tatli < 10000)) {
            System.out.print("Yanlış sayı girdiniz tekrardan deneyiniz");
        }
    }
}