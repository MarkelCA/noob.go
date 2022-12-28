import java.math.BigInteger;

class Factorial {
    public static void main(String[] args) {
        //System.out.println("Java");
        Factorial.factorialHavingLargeResult(2000000);
    }

    public static BigInteger factorialHavingLargeResult(int n) {
        BigInteger result = BigInteger.ONE;
        for (int i = 2; i <= n; i++)
            result = result.multiply(BigInteger.valueOf(i));
        return result;
    }
}
