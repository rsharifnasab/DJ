int [] heap;
int len;
int i;

void heapsort(){
	while(i > 0) {
		pop_max();
	}
}

void add_heap(int x) {
	heap[i] = x;
	relax_up(i);
}

void pop_max() {
	i = i - 1;
	swap(0, i);
	relax(0);
}

void swap(int k, int l) {
	int temp;
	temp = heap[k];
	heap[k] = heap[l];
	heap[l] = temp;
}

void relax_up(int h) {
	int parent;
	if (h == 0) return;
	parent = (h - 1) / 2;
	if (heap[parent] < heap[h]) {
		swap(parent, h);
		relax_up(parent);
	}
}

void relax(int h) {
	int left;
	int right;
	int lindex;
	int rindex;
	int max;
	lindex = 2 * h + 1;
	rindex = 2 * h + 2;
	if (lindex < i) {
		left = heap[lindex];
	} else {
		left = -1;
	}
	if (rindex < i) {
		right = heap[rindex];
	} else {
		right = -1;
	}
	
	max = max(left, right);
	if ((max == 1) && (heap[h] < left)) { 
		swap(h, lindex);
		relax(lindex);
		return;
	}
	if ((max == 2) && (heap[h] < right)) {
		swap(h, rindex);
		relax(rindex);
		return; 
	}

}

int max(int left, int right) {
	if (left > right) return 1;
	return 2;
}


int main() {
	int j;
	len = ReadInteger();
	heap = NewArray(len, int);
	for (i = 0; i < len; i = i + 1) {
		add_heap(ReadInteger());
	}
	heapsort();
	j = 0;
	while (j < len) {
		Print(heap[j]);
		j = j + 1;
	}
}