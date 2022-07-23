int [] distance;
int m;
int [] dest;
int [] src;
int [] weight;
	
void belman_ford() {
	int i;
	for(i = 0;i < m;i = i + 1) {
		relax(i);
	}
}

void relax(int index) {
	if (distance[src[index]] + weight[index] < distance[dest[index]]) {
		distance[dest[index]] = distance[src[index]] + weight[index];
	}
}

int main() {
	int n;
	int i;
	n = ReadInteger();
	m = ReadInteger();
	distance = NewArray(n, int);
	dest = NewArray(m, int);
	src = NewArray(m, int);
	weight = NewArray(m, int);
	for(i = 0;i < m;i = i + 1) {
		Print("Write ", i, "th edge:");
		src[i] = ReadInteger();
		dest[i] = ReadInteger();
		weight[i] = ReadInteger();
	}
	for(i = 0;i < n;i = i + 1) {
		if (i == 0) distance[i] = 0;
		else distance[i] = 1000000;
	}
	for(i = 0;i < n;i = i + 1) {
		belman_ford();
	}
	for(i = 0;i < n; i = i + 1) 
		Print("Distance of ", i, "th node: ", distance[i]);
}