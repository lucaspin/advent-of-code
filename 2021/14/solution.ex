[templates, rules] = File.read!("./input.txt")
  |> String.split("\n\n")
  |> Enum.map(fn line -> String.split(line, "\n") end)

[template] = templates
rule_map =
  rules
  |> Enum.reduce(%{}, fn rule, map ->
    [key, value] = String.split(rule, " -> ")
    Map.put(map, key, value)
  end)

# TODO: this is not good enough for part 2
final_polymer = Enum.reduce(1..10, String.split(template, "", trim: true), fn _step, polymer ->
  transformed =
    polymer
    |> Enum.chunk_every(2, 1, :discard)
    |> Enum.map(fn [left, right] ->
      {:ok, value} = Map.fetch(rule_map, left <> right)
      [left, value, right]
    end)

  List.flatten(
    for {chunk, index} <- Enum.with_index(transformed) do
      [_first | rest] = chunk
      if index == 0, do: chunk, else: rest
    end
  )
end)

frequencies =
  final_polymer
  |> Enum.frequencies()

sorted =
  frequencies
  |> Enum.map(fn {_k, v} -> v end)
  |> Enum.sort()

[min | _rest] = sorted
max = sorted |> List.last()
IO.inspect(min)
IO.inspect(max)
IO.inspect(max - min)
