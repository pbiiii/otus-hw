package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = false

var testCases = []struct {
	text                 string
	expected             []string
	expectedWithAsterisk []string
}{
	{
		text: `Как видите, он  спускается  по  лестнице  вслед  за  своим
			другом   Кристофером   Робином,   головой   вниз,  пересчитывая
			ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
			сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
				кажется, что можно бы найти какой-то другой способ, если бы  он
			только   мог   на  минутку  перестать  бумкать  и  как  следует
			сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
				Как бы то ни было, вот он уже спустился  и  готов  с  вами
			познакомиться.
			- Винни-Пух. Очень приятно!
				Вас,  вероятно,  удивляет, почему его так странно зовут, а
			если вы знаете английский, то вы удивитесь еще больше.
				Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
			вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
			лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
			очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
			громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
			можешь  сделать вид, что ты просто понарошку стрелял; а если ты
			звал его тихо, то все подумают, что ты  просто  подул  себе  на
			нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
			Робин решил отдать его своему медвежонку, чтобы оно не  пропало
			зря.
				А  Винни - так звали самую лучшую, самую добрую медведицу
			в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
			Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
			честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
			знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
			забыл.
				Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
				Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
			иногда,  особенно  когда  папа  дома,  он больше любит тихонько
			посидеть у огня и послушать какую-нибудь интересную сказку.
				В этот вечер...`,
		expected: []string{
			"он",        // 8
			"а",         // 6
			"и",         // 6
			"ты",        // 5
			"что",       // 5
			"-",         // 4
			"Кристофер", // 4
			"если",      // 4
			"не",        // 4
			"то",        // 4
		},
		expectedWithAsterisk: []string{
			"а",         // 8
			"он",        // 8
			"и",         // 6
			"ты",        // 5
			"что",       // 5
			"в",         // 4
			"его",       // 4
			"если",      // 4
			"кристофер", // 4
			"не",        // 4
		},
	},
	{
		text: `Lorem ipsum dolor sit amet, consectetur adipiscing elit.
			Praesent congue in est sed semper. Phasellus magna elit, dapibus ac est ac,
			eleifend convallis dolor. Aenean nunc leo, faucibus vitae nisi in,
			malesuada scelerisque eros. Suspendisse potenti. Vestibulum eu eros felis.
			Morbi vehicula magna purus, non imperdiet neque cursus id.
			Proin condimentum mollis mi. Vestibulum lobortis libero eu mauris euismod,
			non vehicula dui mollis. Praesent gravida gravida erat, in bibendum nisi fringilla ac.
			Interdum et malesuada fames ac ante ipsum primis in faucibus. Sed aliquam,
			purus eget pretium mattis, orci elit suscipit dui, quis suscipit tortor ex non lorem.
			In id massa eu erat imperdiet lacinia ac id elit. Quisque elementum pharetra orci,
			nec vestibulum justo viverra nec. Donec nec semper felis.
			
			Sed arcu metus, condimentum ac vehicula id, euismod sed libero.
			Pellentesque at rhoncus lorem, eget cursus mi. Maecenas congue diam id ante imperdiet euismod.
			Ut lorem dui, accumsan a bibendum interdum, facilisis sed nisl. Duis turpis nisl,
			semper in augue a, commodo eleifend velit. Nam eleifend euismod euismod.
			Praesent a dictum nisi, sit amet bibendum elit. In ac lacinia turpis.
			Sed sapien quam, tempor ac suscipit scelerisque, venenatis ac nisi. Donec ut odio quam.
			Sed sollicitudin dolor in enim malesuada, eget tincidunt libero bibendum.
			Donec eget finibus urna. Curabitur eleifend, ipsum et condimentum laoreet,
			nulla ligula sagittis diam, a sollicitudin nulla elit nec eros.
			Duis sodales at nisi id hendrerit.`,
		expected: []string{
			"ac",          // 7
			"in",          // 5
			"Sed",         // 4
			"eget",        // 4
			"id",          // 4
			"Donec",       // 3
			"Praesent",    // 3
			"a",           // 3
			"bibendum",    // 3
			"condimentum", // 3
		},
	},
}

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	for _, testCase := range testCases {
		t.Run("positive test", func(t *testing.T) {
			if taskWithAsteriskIsCompleted {
				require.Equal(t, testCase.expectedWithAsterisk, Top10(testCase.text))
			} else {
				require.Equal(t, testCase.expected, Top10(testCase.text))
			}
		})
	}
}
