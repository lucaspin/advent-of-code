defmodule PathFinder do
  def load do
    File.read!("./input3.txt")
    |> String.split("\n", trim: true)
    |> Enum.map(fn line -> String.split(line, "-") end)
    |> Enum.reduce(%{}, fn [from, to], graph ->
      updated = Map.update(graph, from, [to], fn cur -> [to | cur] end)
      Map.update(updated, to, [from], fn cur -> [from | cur] end)
    end)
  end

  def recur(_graph, _visited, cur, dest) when cur == dest, do: 1
  def recur(graph, visited, cur, dest) do
    Enum.reduce(graph[cur], 0, fn node, count ->
      if is_uppercase?(node) || !Enum.member?(visited, node) do
        count + recur(graph, [cur | visited], node, dest)
      else
        count
      end
    end)
  end

  def recur2(_graph, _visited, cur, dest) when cur == dest, do: 1
  def recur2(graph, visited, cur, dest) do
    Enum.reduce(graph[cur], 0, fn node, count ->
      if can_visit?(node, [cur | visited]) do
        count + recur2(graph, [cur | visited], node, dest)
      else
        count
      end
    end)
  end

  def can_visit?("start", _visited), do: false
  def can_visit?("end", _visited), do: true
  def can_visit?(node, visited) do
    cond do
      is_uppercase?(node) -> true
      !Enum.member?(visited, node) -> true
      true ->
        visited
        |> Enum.filter(fn node -> !is_uppercase?(node) end)
        |> Enum.frequencies()
        |> Enum.all?(fn {_node, count} -> count == 1 end)
    end
  end

  def is_uppercase?(string) do
    string == String.upcase(string)
  end
end

# Part 1
# PathFinder.load()
# |> PathFinder.recur([], "start", "end")
# |> IO.inspect()

# Part 2
PathFinder.load()
|> PathFinder.recur2([], "start", "end")
|> IO.inspect()
