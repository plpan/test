namespace go thrift.example

struct Numbers {
    1: required i32 a;
    2: required i32 b;
}

service Adder {
	i32 add(1:i32 a, 2:i32 b);
    Numbers addNumber(1:Numbers a, 2:Numbers b);
}
