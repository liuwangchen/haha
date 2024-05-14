package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func findFirstValue(queryKey string, raw json.RawMessage) (string, error) {
	var asMap map[string]*json.RawMessage
	if err := json.Unmarshal(raw, &asMap); err == nil {
		for key, val := range asMap {
			if strings.EqualFold(key, queryKey) {
				return string(*val), nil // 直接返回原始 JSON 字符串
			}

			value, err := findFirstValue(queryKey, *val)
			if err != nil {
				return "", err
			}
			if value != "" {
				return value, nil
			}
		}
		return "", nil
	}

	var asSlice []*json.RawMessage
	if err := json.Unmarshal(raw, &asSlice); err == nil {
		for _, val := range asSlice {
			value, err := findFirstValue(queryKey, *val)
			if err != nil {
				return "", err
			}
			if value != "" {
				return value, nil
			}
		}
	}

	return "", nil
}

func main() {
	jsonData := `{"Info":{"HallConfigID":"FEATURE_HALL_GRASSY_GROOMERS_1", "RoundID":"FEATURE_HALL_GRASSY_GROOMERS_1.2gCeBlM51IZjLl5WRWM8f2cijPi.85", "Cards":[{"Card":{"Squares":[{"Num":15}, {"Num":14}, {"Num":10, "Atts":[{"Type":5, "Count":1}]}, {"Num":2, "Mod":{"Type":10, "Count":1}, "Atts":[{"Type":5, "Count":1}]}, {"Num":1}, {"Num":27}, {"Num":19}, {"Num":30, "Atts":[{"Type":3, "Count":1}]}, {"Num":26, "Mod":{"Type":8, "Count":1}}, {"Num":21, "Atts":[{"Type":4, "Count":1}]}, {"Num":38, "Mod":{"Type":11, "Count":1}}, {"Num":40, "Atts":[{"Type":4, "Count":1}]}, {"Num":33}, {"Num":36}, {"Num":32, "Atts":[{"Type":1, "Count":1}]}, {"Num":47}, {"Num":50, "Atts":[{"Type":2, "Count":1}]}, {"Num":58, "Atts":[{"Type":2, "Count":1}]}, {"Num":60, "Mod":{"Type":8, "Count":1}, "Atts":[{"Type":1, "Count":1}]}, {"Num":59, "Atts":[{"Type":3, "Count":1}]}, {"Num":70}, {"Num":75}, {"Num":71, "Mod":{"Type":8, "Count":1}}, {"Num":68, "Atts":[{"Type":4, "Count":1}]}, {"Num":74, "Atts":[{"Type":2, "Count":1}]}], "IsCollectionCard":true, "BombIndices":[19, 13, 16], "DIndices":[21, 17, 1, 24, 2, 9, 19, 5], "AIndices":[14, 21, 20, 4, 22, 2, 13], "BIndices":[11, 15, 8, 24, 21, 14, 16, 17, 13, 19, 18, 12], "EIndices":[6, 3, 7, 10, 0, 23], "Type":"2"}, "Daubs":[15, 14, 10, 2, 1, 27, 19, 30, 26, 21, 38, 40, 33, 36, 32, 47, 50, 58, 60, 59, 70, 75, 71, 68, 74], "BingoLevel":1, "BingoRank":31, "BingoLevelReward":{"Resources":[{"ID":"XP", "Amount":5000}, {"ID":"TICKET", "Amount":10}]}, "DaubReward":{"Resources":[{"ID":"ACP_GAMEPLAY_1", "Amount":1}, {"ID":"XP", "Amount":1430}, {"ID":"COIN", "Amount":9}]}, "CollectionReward":{"Resources":[{"ID":"F_1_1_2", "Amount":1}]}, "CampaignReward":[{"CampaignID":"MERGE_10", "Reward":{"Resources":[{"ID":"MERGE_ENERGY", "Amount":1}]}}, {"CampaignID":"MANIA_38", "Reward":{"Resources":[{"ID":"SHIP_WHEEL", "Amount":10}]}}]}, {"Card":{"Squares":[{"Num":5, "Mod":{"Type":8, "Count":1}, "Atts":[{"Type":2, "Count":1}]}, {"Num":14, "Atts":[{"Type":2, "Count":1}]}, {"Num":9}, {"Num":6}, {"Num":1, "Atts":[{"Type":1, "Count":1}]}, {"Num":24}, {"Num":27, "Atts":[{"Type":5, "Count":1}]}, {"Num":18, "Mod":{"Type":4, "Count":1}}, {"Num":28}, {"Num":16}, {"Num":39, "Atts":[{"Type":3, "Count":1}]}, {"Num":35}, {"Num":36, "Atts":[{"Type":5, "Count":1}]}, {"Num":31}, {"Num":45, "Atts":[{"Type":1, "Count":1}]}, {"Num":58, "Atts":[{"Type":4, "Count":1}]}, {"Num":53}, {"Num":56, "Atts":[{"Type":2, "Count":1}]}, {"Num":51, "Mod":{"Type":8, "Count":1}, "Atts":[{"Type":3, "Count":1}]}, {"Num":49}, {"Num":65}, {"Num":66, "Atts":[{"Type":4, "Count":1}]}, {"Num":62}, {"Num":64, "Atts":[{"Type":4, "Count":1}]}, {"Num":72}], "BombIndices":[21, 21, 3], "DIndices":[0, 6, 4, 20, 23, 9, 21, 14, 7, 15, 1], "AIndices":[0, 18, 7, 23, 22, 21, 12, 2, 5], "BIndices":[22, 3, 21, 8, 2, 16, 4, 20], "EIndices":[19, 11, 10, 17, 24, 13], "Type":"1"}, "Daubs":[14, 24, 27, 18, 28, 16, 45, 58, 53, 56, 51, 49, 66], "DaubReward":{"Resources":[{"ID":"COIN", "Amount":3}, {"ID":"XP", "Amount":780}]}}], "DoubleXP":true, "Rankings":[{"ID":"robot-user-4895", "Rank":1, "Score":"2100", "BetID":"BET_2"}, {"ID":"robot-user-6552", "Rank":2, "Score":"2091", "BetID":"BET_4"}, {"ID":"robot-user-5217", "Rank":3, "Score":"2089", "BetID":"BET_3"}, {"ID":"robot-user-6420", "Rank":4, "Score":"1799", "BetID":"BET_5"}, {"ID":"robot-user-7134", "Rank":5, "Score":"1659", "BetID":"BET_5"}, {"ID":"robot-user-7139", "Rank":6, "Score":"1575", "BetID":"BET_5"}, {"ID":"robot-user-7056", "Rank":7, "Score":"1573", "BetID":"BET_3"}, {"ID":"robot-user-7000", "Rank":8, "Score":"1511", "BetID":"BET_5"}, {"ID":"robot-user-5250", "Rank":9, "Score":"1077", "BetID":"BET_5"}, {"ID":"robot-user-7110", "Rank":10, "Score":"1051", "BetID":"BET_4"}, {"ID":"robot-user-7052", "Rank":11, "Score":"975", "BetID":"BET_4"}, {"ID":"robot-user-7112", "Rank":12, "Score":"975", "BetID":"BET_3"}, {"ID":"robot-user-7105", "Rank":13, "Score":"965", "BetID":"BET_2"}, {"ID":"2dk5B0aD4BRzmIxONr1R1NbYz3Z", "Rank":14, "Score":"954", "BetID":"BET_FEATURE_GROUP_1_2"}, {"ID":"robot-user-4824", "Rank":15, "Score":"953", "BetID":"BET_4"}, {"ID":"robot-user-5243", "Rank":16, "Score":"940", "BetID":"BET_3"}, {"ID":"robot-user-6292", "Rank":17, "Score":"919", "BetID":"BET_4"}, {"ID":"robot-user-6285", "Rank":18, "Score":"862", "BetID":"BET_2"}, {"ID":"robot-user-7014", "Rank":19, "Score":"839", "BetID":"BET_5"}, {"ID":"robot-user-7128", "Rank":20, "Score":"826", "BetID":"BET_3"}, {"ID":"2g80826EasxW53dqaAJAwy1k9gY", "Rank":21, "Score":"821", "BetID":"BET_FEATURE_GROUP_1_2"}, {"ID":"robot-user-7129", "Rank":22, "Score":"813", "BetID":"BET_4"}, {"ID":"robot-user-7019", "Rank":23, "Score":"695", "BetID":"BET_3"}, {"ID":"robot-user-5504", "Rank":24, "Score":"543", "BetID":"BET_2"}, {"ID":"robot-user-7080", "Rank":25, "Score":"500", "BetID":"BET_4"}, {"ID":"robot-user-7005", "Rank":26, "Score":"488", "BetID":"BET_2"}, {"ID":"robot-user-7187", "Rank":27, "Score":"278", "BetID":"BET_3"}, {"ID":"robot-user-7135", "Rank":28, "Score":"151", "BetID":"BET_5"}], "LastRewardRank":9, "PerfectDaubReward":{"Resources":[{"ID":"XP", "Amount":1000}]}, "TicketCost":20, "BetID":"BET_FEATURE_GROUP_1_2", "LuckyPatternWheel":{}, "HallID":"2gCeBlM51IZjLl5WRWM8f2cijPi"}}`
	queryKey := "Rankings"

	raw := json.RawMessage(jsonData)

	value, err := findFirstValue(queryKey, raw)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Value: %s\n", value)
}
