class QueueItem {
  int data;
  QueueItem next;
  QueueItem prev;

  void Init(int data, QueueItem next, QueueItem prev) {
    this.data = data;
    this.next = next;
    next.prev = this;
    this.prev = prev;
    prev.next = this;
  }
  int GetData() {
    return this.data;
  }
  QueueItem GetNext() { return next;}
  QueueItem GetPrev() { return prev;}
  void SetNext(QueueItem n) { next = n;}
  void SetPrev(QueueItem p) { prev = p;}
}


class Queue {
  QueueItem head;

  void Init() {
    this.head = new QueueItem;
    this.head.Init(0, this.head, this.head);
  }
  void EnQueue(int i) {
    QueueItem temp;
    temp = new QueueItem;
    temp.Init(i, this.head.GetNext(), this.head);
  }
  int DeQueue() {
    int val;
    if (this.head.GetPrev() == this.head) {
      Print("Queue Is Empty");
      return 0;
    } else {
      QueueItem temp;
      temp = this.head.GetPrev();
      val = temp.GetData();
      temp.GetPrev().SetNext(temp.GetNext());
      temp.GetNext().SetPrev(temp.GetPrev());
    }
    return val;
  }
}

void main() {
  Queue q;
  int i;
  q = new Queue;
  q.Init();
  for (i = 0; i != 10; i = i + 1)
    q.EnQueue(i);

  for (i = 0; i != 4; i = i + 1)
    Print(q.DeQueue());
  Print("\n");
  for (i = 0; i != 10; i = i + 1)
    q.EnQueue(i);

  for (i = 0; i != 17; i = i +1)
    Print(q.DeQueue());
}