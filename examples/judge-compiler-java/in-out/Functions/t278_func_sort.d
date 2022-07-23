int[] ReadArray()
{
  int i;
  int num;
  int [] arr;
  int numScores;

  Print("How many scores?");
  numScores = ReadInteger();
  arr = NewArray(numScores, int);
  i = 0;
  while (i < arr.length()) {
    Print("Enter next number:");
    num = ReadInteger();
    arr[i] = num;
    i = i + 1;
  }
  return arr;
}

void Sort(int []arr)
{
  int i;
  int j;
  int val;

  i = 1;
  while (i < arr.length()) {
    j = i -1;
    val = arr[i];
    while (j >= 0) {
      if (val >= arr[j]) break;
	arr[j+1] = arr[j];
      j = j -1;
   }
   arr[j+1] = val;
   i = i + 1;
  }
}

void PrintArray(int []arr)
{
  int i;
   i = 0;
   Print("Sorted results:");
   while (i < arr.length()) {
	Print(arr[i]);
	i = i + 1;
  }
}


void main()
{
  int[] arr;

  Print("This program will read in a bunch of numbers and print them");
  Print("back out in sorted order.");
  arr = ReadArray();
  Sort(arr);
  PrintArray(arr);
}