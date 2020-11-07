import java.io.*;
import java.util.StringTokenizer;

public class JobQueue {

    public int numWorkers;
    private int[] jobs;

    private long[] assignedWorker;
    private long[] startTime;

    private FastScanner in;
    private PrintWriter out;

    public static void main(String[] args) throws IOException {
        new JobQueue().solve();
    }

    private void readData() throws IOException {
        numWorkers = in.nextInt();
        int m = in.nextInt();
        jobs = new int[m];
        for (int i = 0; i < m; ++i) {
            jobs[i] = in.nextInt();
        }
    }

    private void writeResponse() {
        for (int i = 0; i < jobs.length; ++i) {
            out.println(assignedWorker[i] + " " + startTime[i]);
        }
    }

    private void assignJobs() {
        // making dynamic arrays for the assigned worker and the start time for the
        // particular jobs
        assignedWorker = new long[jobs.length];
        startTime = new long[jobs.length];

        //instance for instantiating the min heap class 
        minHeapImplement assigningJobs = new minHeapImplement(numWorkers);

        int i = 0;

        // assigning the first batches to the first jobs in a row
        for (i = 0; i < jobs.length && i < numWorkers; i++) {
            assignedWorker[i] = i;
            startTime[i] = 0;
            assigningJobs.insert(jobs[i]);
        }

        //assigning the rest of the jobs
        while (i < jobs.length) {
            long[] newJob = assigningJobs.extractAndInsert(jobs[i]);
            assignedWorker[i] = newJob[0];
            startTime[i] = newJob[1];
            i++;
        }

    }

    public void solve() throws IOException {
        in = new FastScanner();
        out = new PrintWriter(new BufferedOutputStream(System.out));
        readData();
        assignJobs();
        writeResponse();
        out.close();
    }

    //scanner class for fast input from the stdin stream
    static class FastScanner {
        private BufferedReader reader;
        private StringTokenizer tokenizer;

        public FastScanner() {
            reader = new BufferedReader(new InputStreamReader(System.in));
            tokenizer = null;
        }

        public String next() throws IOException {
            while (tokenizer == null || !tokenizer.hasMoreTokens()) {
                tokenizer = new StringTokenizer(reader.readLine());
            }
            return tokenizer.nextToken();
        }

        public int nextInt() throws IOException {
            return Integer.parseInt(next());
        }
    }

    public class minHeapImplement {

        private long[] minHeap;
        private int noOfElems;

        // extra layer
        private long[] assignedT;
        private long[] done;
        
        //constructor for our class
        minHeapImplement(int treeHeight) {
            minHeap = new long[treeHeight];
            assignedT = new long[treeHeight];
            for (int i = 0; i < treeHeight; i++) {
                assignedT[i] = i;
            }
            noOfElems = 0;

            done = new long[2];
        }

        //for swapping the two elements in an array
        private void swap(long[] array, int a, int b) {
            long store = array[a];
            array[a] = array[b];
            array[b] = store;
        }

        public void siftUp(int index) {
            if (index == 0) {
                return;
            }

            boolean first = minHeap[index] < minHeap[(index - 1) / 2];
            boolean second = minHeap[index] == minHeap[(index - 1) / 2];
            boolean third = assignedT[index] < assignedT[(index - 1) / 2];

            if (first || (second && third)) {
                swap(minHeap, index, (index - 1) / 2);
                swap(assignedT, index, (index - 1) / 2);
                siftUp((index - 1) / 2);
            }

        }

        public void siftDown(int index) {

            //element with the smallestIndex
            int smallestIndexElem = index;

            boolean first , second , combined;

            if (2 * index + 1 < noOfElems) {

                //additional booleans second and third for the jobs to work properly
                first = minHeap[smallestIndexElem] > minHeap[2 * index + 1];
                second = minHeap[smallestIndexElem] == minHeap[2*index + 1] && assignedT[smallestIndexElem] > assignedT[2*index + 1];
                combined = first || second;

                if(combined){
                    smallestIndexElem = 2 * index + 1;    
                }
            }

            if (2 * index + 2 < noOfElems) {

                first = minHeap[smallestIndexElem] > minHeap[2 * index + 2];
                second = minHeap[smallestIndexElem] == minHeap[2*index + 2] && assignedT[smallestIndexElem] > assignedT[2*index + 2];
                combined = first || second;

                if(combined){
                    smallestIndexElem = 2 * index + 2;    
                }
            }

            if (index == smallestIndexElem) {
                return;
            }

            //swapping elements from both the minHeap and the assigned thread section
            swap(minHeap, index, smallestIndexElem);
            swap(assignedT, index, smallestIndexElem);

            //sifting down the said element
            siftDown(smallestIndexElem);
        }

        public long[] extractAndInsert(int elem) {

            //using Done array for returning the thread freed and the job done
            done[0] = assignedT[0];
            done[1] = minHeap[0];

            //assigning to the next job to the thread that was freed
            minHeap[0] = elem + minHeap[0];

            //sifting down it's priority
            siftDown(0);

            //returning the done[] to the main class
            return done;
        }

        public void insert(int elem) {
            //inserting at the last
            minHeap[noOfElems] = elem;

            //sifting the said element up if needed
            siftUp(noOfElems);

            //increasing the number of elements by one
            noOfElems++;
        }

    }

}
