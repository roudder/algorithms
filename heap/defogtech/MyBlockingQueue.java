package algorithms.heap.defogtech;

import java.util.LinkedList;
import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.ReentrantLock;

public class MyBlockingQueue<E> {
    private ReentrantLock lock = new ReentrantLock(true);
    private Condition empty;
    private Condition notEmpty;
    private Condition full;
    private Condition notFull;
    private LinkedList<E> queue;
    private int size;

    MyBlockingQueue(Integer size) {
        queue = new LinkedList<>();
        notEmpty = lock.newCondition();
        notFull = lock.newCondition();
        this.size = size;
    }

    void put(E e) {
        lock.lock();
        try {
            while (queue.size() == size) {
                notFull.await();
            }
            queue.push(e);
            notEmpty.signalAll();
        } catch (InterruptedException ex) {
            ex.printStackTrace();
        } finally {
            lock.unlock();
        }
    }


    void get() {
        lock.lock();
        try {
            while (queue.size() == 0) {
                notEmpty.await();
            }
            E e = queue.pop();
            notFull.signalAll();
        } catch (InterruptedException ex) {
            ex.printStackTrace();
        } finally {
            lock.unlock();
        }
    }
}


