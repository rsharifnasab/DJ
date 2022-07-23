class table {
	int [][] table;
	string [] chars;
	
	void init() {
		int i;
		int j;
		this.table = NewArray(3, int []);
		for (i = 0;i < 3;i = i + 1) {
			this.table[i] = NewArray(3, int);
		}
		for (i = 0;i < 3;i = i + 1) 
			for (j = 0;j < 3;j = j + 1)
				table[i][j] = 0;
		
		this.chars = NewArray(3, string);
		this.chars[0] = "E";
		this.chars[1] = "X";
		this.chars[2] = "O";
	}
	
	bool set(int x, int y, int player) {
		if (this.table[x][y] != 0) {
			Print("Illegal point!");
			return false;
		}
		this.table[x][y] = player;
		return true;
	}
	
	int winner() {
		int i;
		int player;
		i = 0;
		while (i < 3) {
			player = this.table[i][0];
			if (this.table[i][0] == this.table[i][1] && this.table[i][1] == this.table[i][2] && player != 0)
				return player;
			player = this.table[0][i];
			if (this.table[0][i] == this.table[1][i] && this.table[1][i] == this.table[2][i] && player != 0)
				return player;
			i = i + 1;
		}
		if (this.table[0][0] == this.table[1][1] && this.table[1][1] == this.table[2][2] && this.table[0][0] != 0)
			return this.table[0][0];
		if (this.table[0][2] == this.table[1][1] && this.table[1][1] == this.table[2][0] && this.table[2][0] != 0)
			return this.table[0][2];
		return 0;
	}
	
	bool is_full() {
		int i;
		int j;
		for (i = 0;i < 3;i = i + 1) 
			for (j = 0;j < 3;j = j + 1)
				if (table[i][j] == 0)
					return false;
		return true;
	}
	
	void print() {
		int i;
		for (i = 0;i < 3;i = i + 1) {
			Print(chars[table[i][0]], "|", chars[table[i][1]], "|", chars[table[i][2]]);
		}
	}
}

class dooz {
	table table;
	int nobat;
	
	void init() {
		table = new table;
		table.init();
		this.nobat = 1;
	}
	
	bool round(int x, int y) {
		bool ret;
		int win;
		ret = table.set(x, y, this.nobat);
		if (!ret) {
			Print("Try another point!");
			return false;
		}
		win = table.winner();
		if (win != 0) {
			Print("Player ", win, " winnes!");
			return true;
		}
		if (table.is_full()) {
			Print("Draw!");
			return true;
		}
		this.nobat = 3 - this.nobat;
		table.print();
		Print("player", nobat, " choose your point:");
		return false;
	}
}

int main() {
	dooz game;
	int x;
	int y;
	game = new dooz;
	game.init();
	Print("Player1 choose your point:");
	x = ReadInteger();
	y = ReadInteger();
	while(!game.round(x - 1, y - 1)) {
		x = ReadInteger();
		y = ReadInteger();
	}
}