defmodule Polymer do
  def solve(steps) do
    {template, rules, pair_frequencies} = parse()
    recur(template, rules, pair_frequencies, steps)
  end

  def recur(template, _rules, pair_frequencies, 0), do: result(template, pair_frequencies)
  def recur(template, rules, pair_frequencies, steps_remaining) do
    new_frequencies = Enum.reduce(pair_frequencies, Map.new(), fn {pair, count}, frequencies ->
      new_letter = rules[pair]
      [left, right] = pair |> String.graphemes()
      Map.update(frequencies, left <> new_letter, count, fn cur -> cur + count end)
      |> Map.update(new_letter <> right, count, fn cur -> cur + count end)
    end)

    recur(template, rules, new_frequencies, steps_remaining - 1)
  end

  def result(template, pair_frequencies) do
    last_letter = template |> String.graphemes() |> List.last()
    initial = Map.put(Map.new(), last_letter, 1)
    Enum.reduce(pair_frequencies, initial, fn {pair, count}, letter_count_map ->
      [left, _right] = pair |> String.graphemes()
      Map.update(letter_count_map, left, count, fn cur -> cur + count end)
    end)
  end

  def parse do
    [[template], rules] =
      File.read!("./input.txt")
      |> String.split("\n\n")
      |> Enum.map(fn line -> String.split(line, "\n") end)

    rules =
      rules
      |> Enum.reduce(Map.new(), fn rule, frequencies ->
        [pair, value] = rule |> String.split(" -> ")
        Map.put(frequencies, pair, value)
      end)

    pair_frequencies =
      String.split(template, "", trim: true)
      |> Enum.chunk_every(2, 1, :discard)
      |> Enum.reduce(Map.new(), fn [left, right], frequencies ->
        Map.update(frequencies, left <> right, 1, fn cur -> cur + 1 end)
      end)

    {template, rules, pair_frequencies}
  end
end

Polymer.solve(40)
|> Map.values()
|> Enum.min_max()
|> then(fn {min, max} -> max - min end)
|> IO.inspect()
