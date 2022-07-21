
	int[] array;
	int[] tempArray;

	void merge(int s1, int s2, int e1, int e2)
	{
		int i;
		int j;
		int k;
		int curlen;
		int it;
		int newIt;
		int loopI;

		j = s2;
		k = 0;
		curlen = e2 - s1 + 1;
		it = s1;
		loopI = 11;

		for (i = s1; i<e1+1; i=i+1){
			if(j > e2){
				loopI = i;
				break;
			}
			if(array[i] < array[j]){
				tempArray[k] = array[i];
				k = k + 1;
			}
			else{
				tempArray[k] = array[j];
				k = k + 1;
				j = j + 1;
				i = i - 1;
			}
		}

		if(loopI <= e1)
		{
			for (newIt = loopI; newIt<e1+1; newIt = newIt + 1){
				tempArray[k] = array[newIt];
				k = k + 1;
			}
		}

		if(j <= e2)
		{
			for (newIt = j; newIt<e2+1; newIt = newIt + 1){
				tempArray[k] = array[newIt];
				k = k + 1;
			}
		}

		for (k = 0; k<curlen; k=k+1) {
			array[it] = tempArray[k];
			it = it + 1;
		}

	}

	void mergesort(int start, int end){

		if(start < end){
			int mid;
			mid = start + (end - start)/2;
			mergesort(start,mid);
			mergesort(mid+1,end);
			merge(start,mid+1,mid,end);

		}
	}

	void main(){
         int i;

		array = NewArray(10, int);
		tempArray = NewArray(10, int);

		for (i = 0; i<10; i=i+1){
			array[i] = 10 - i;
		}

		mergesort(0,9);


		for (i = 0;  i<10; i=i+1) {
			Print(array[i]);
		}
	}