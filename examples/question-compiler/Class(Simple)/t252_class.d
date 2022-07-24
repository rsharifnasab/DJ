class dynamic_array {
	int [] arr;
	int len;
	
	void init() {
		arr = NewArray(3, int);
		len = 0;
	}
	
	void append(int m) {
		if (len == arr.length()) {
			int i;
			int [] temp;
			temp = arr;
			arr = NewArray(2 * arr.length(), int);
			for(i = 0; i < len; i = i + 1) {
				arr[i] = temp[i];
			}
		}
		arr[len] = m;
		len = len + 1;
	}
	
	int pop() {
		int ret;
		if (len == 0) {
			return 0;
		}
		len = len - 1;
		ret = arr[len];
		if (3 * len < arr.length()) {
			int i;
			int [] temp;
			temp = arr;
			arr = NewArray(arr.length() / 2 + 1, int);
			for (i = 0;i < len;i = i + 1) {
				arr[i] = temp[i];
			}
		}
		return ret;
	}
	
	void print() {
		int i;
		for(i = 0;i < len;i = i + 1) {
			Print(i, "th element of array : ", arr[i]);
		}
	}
}


int main() {
	int j;
	dynamic_array t;
	t = new dynamic_array;
	t.init();
	for (j = 0;j < 40;j = j + 1) {
		if (j % 8 != 7) {
			t.append(j);
		} else {
			t.pop();
		}
		Print(j," elements of array : ");
		t.print();
	}
	
}
