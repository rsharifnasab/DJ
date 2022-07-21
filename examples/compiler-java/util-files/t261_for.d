int main(){
     int i;
	int j;
     i = 0;
    for(; i < 3; i = i+1){
		for(j=i; j >= 0; j = j-1){
			Print(i, j);
		}
	}
}