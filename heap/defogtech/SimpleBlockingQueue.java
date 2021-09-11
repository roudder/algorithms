package algorithms.heap.defogtech;

import java.util.Random;
import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.BlockingQueue;

public class SimpleBlockingQueue {

    public static void main(String[] args) {
        BlockingQueue<Integer> queue = new ArrayBlockingQueue<>(100);
        Runnable producer = () -> {
            while (true) {
                try {
                    Thread.currentThread().sleep(2000);
                    Integer random = new Random().nextInt(100);
                    queue.put(random);
                    System.out.println("remaining cap. after put " + queue.remainingCapacity());
                    System.out.println("put " + random);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        };
        new Thread(producer).start();
        new Thread(producer).start();

        Runnable consumer = () -> {
            while (true) {
                try {
                    Thread.currentThread().sleep(3000);
                    Integer obj = queue.take();
                    System.out.println("remaining cap. after take " + queue.remainingCapacity());
                    System.out.println("consume " + obj);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        };
        new Thread(consumer).start();
        new Thread(consumer).start();
    }
}
