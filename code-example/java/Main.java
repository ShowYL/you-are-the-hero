import java.math.BigInteger;
import java.time.Duration;
import java.time.Instant;
import java.util.ArrayList;

class Main {

    static final int max = 50000;
    static final int min = 20000;
    static final int amount = 1000;

    void main() {
        IO.println("Starting threads.");

        var start = Instant.now();

        // parallelism();
        // sequential();

        var end = Instant.now();

        IO.println(
            "All threads finished in " +
                Duration.between(start, end).toMillis() +
                " ms"
        );
    }

    void sequential() {
        IO.println("Sequential");
        for (int i = 0; i < amount; i++) {
            int randomNum = min + (int) (Math.random() * (max - min + 1));

            var compute = new Compute(randomNum);
            compute.run();
        }
    }

    void parallelism() {
        IO.println("parallelism");
        var threads = new ArrayList<Thread>();
        for (int i = 0; i < amount; i++) {
            int randomNum = min + (int) (Math.random() * (max - min + 1));

            var compute = new Compute(randomNum);
            var worker = new Thread(compute);
            worker.start();
            threads.add(worker);
        }

        for (var thread : threads) {
            try {
                thread.join();
            } catch (InterruptedException e) {
                IO.println("Error while waiting for the thread: " + e);
            }
        }
    }

    class Compute implements Runnable {

        private int number;

        Compute(int n) {
            this.number = n;
        }

        @Override
        public void run() {
            BigInteger prev = BigInteger.valueOf(0);
            BigInteger prevprev = BigInteger.valueOf(1);
            int idx = 0;

            while (idx != number) {
                BigInteger current = prev.add(prevprev);
                prevprev = prev;
                prev = current;
                idx++;
            }

            // IO.println("Fibonacci for " + number + ": " + prev);
        }
    }
}
