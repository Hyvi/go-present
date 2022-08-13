% Understanding actor concurrency, Part1: Actors in Erlang, https://www.infoworld.com/article/2077999/understanding-actor-concurrency--part-1--actors-in-erlang.html
-module(temperature_rpc).
-export([temperatureConverter/0, start/0, convert/2]).

start() -> 
	spawn(fun() -> temperatureConverter() end).

temperatureConverter() ->
	receive
		% words starting with lowercase letters are atoms -- fixed symbols 
		% that always represent a constant, number, or string of the same name, 
		% even across machines.
		{From, {toF, C}} ->
			From ! {self(), 32+C*9/5},
			temperatureConverter();
		{From, {toC, F}} -> 
			From ! {self(), (F-32)*5/9},
			temperatureConverter();
		{From, stop} ->
			From ! {self(), "Stopping"}, 
			io:format("Stopping~n");
		{From, Other} -> 
			From ! {self(), "Unkown"},
			io:format("Unkown: ~p~n", [Other]),
			temperatureConverter()
	end.

convert(Pid, Request) ->
	Pid ! {self(), Request}, 
	receive
		{Pid, Response} -> Response
	end.
