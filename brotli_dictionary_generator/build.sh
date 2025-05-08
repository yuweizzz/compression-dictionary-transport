cd libdivsufsort/
cmake . -DBUILD_SHARED_LIBS=OFF
make
cd ..
g++ -c brotli/research/deorummolae.cc -I esaxx/
g++ -c brotli/research/durchschlag.cc -I libdivsufsort/include
g++ -c brotli/research/sieve.cc
g++ -c brotli/research/dictionary_generator.cc
g++ -o brotli_dictionary_generator *.o -L libdivsufsort/lib/ -ldivsufsort -static
