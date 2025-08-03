package main

import (
	"fmt"
	"image/color"
	"net/url"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/rand"
)

func main() {
	a := app.New()
	w := a.NewWindow("Atmosphere")
	w.Resize(fyne.NewSize(800, 716))
	w.CenterOnScreen()

	color_hi := color.RGBA{9, 240, 203, 255}
	hi1 := canvas.NewText("Это правда ты? Привет! А скорее всего, доброй ночи. 💤", color_hi)
	hi2 := canvas.NewText("Похоже, ты хочешь уснуть, но сон всё никак не приходит.", color_hi)
	hi3 := canvas.NewText("Возможно, особая атмосфера поможет тебе погрузиться в мир сновидений...", color_hi)

	cont_hi := container.NewVBox(
		container.NewCenter(
			hi1,
		),
		container.NewCenter(
			hi2,
		),
		container.NewCenter(
			hi3,
		),
	)

	label_Weather := widget.NewLabel("")

	sel := widget.NewSelect(
		[]string{
			"Дождь🌧️",
			"Снег❄️",
			"Грозу🌩️",
			"Ветер🌀",
			"Ураган🌪️",
			"Лунный прибой🌙",
		},
		func(s string) {
			fmt.Printf("Выбрано - %s\n", s)

			label_Weather.SetText(s)
		},
	)
	sel.PlaceHolder = "Какую погоду ты хочешь? Вообрази, расслабься"

	img_mask, _ := fyne.LoadResourceFromPath("./icons/mask.png")
	img_girl := canvas.NewImageFromFile("girl.jpg")
	img_girl.FillMode = canvas.ImageFillOriginal

	btn_Weather_see := widget.NewButtonWithIcon("Хочу увидеть", img_mask, func() {
		label_Weather.SetText("Вы выбрали: " + sel.Selected)
		var url_R *url.URL

		if sel.Selected == "Дождь🌧️" {
			url_R, _ = url.Parse("https://www.youtube.com/watch?v=Zhq_ThHG2gA")
		}

		if sel.Selected == "Снег❄️" {
			url_R, _ = url.Parse("https://www.youtube.com/watch?v=ydaj81_E7LU")
		}

		if sel.Selected == "Грозу🌩️" {
			url_R, _ = url.Parse("https://www.youtube.com/watch?v=umKLBM288G0")
		}

		if sel.Selected == "Ветер🌀" {
			url_R, _ = url.Parse("https://www.youtube.com/watch?v=qMLi6S8YwxI&list=PLESKtX4D9WC85WYsnkoiGPJk_m5-sHqxB&index=11")
		}

		if sel.Selected == "Ураган🌪️" {
			url_R, _ = url.Parse("https://www.youtube.com/watch?v=wZ9Ga8V1q_4")
		}

		if url_R == nil {
			url_R, _ = url.Parse("https://ru.pinterest.com/pin/44332377578601033/")
		}

		if sel.Selected == "Лунный прибой🌙" {
			url_R, _ = url.Parse("https://www.youtube.com/watch?v=_Al8valACfA")
		}

		err := fyne.CurrentApp().OpenURL(url_R)
		if err != nil {
			fyne.LogError("Ошибка", err)
		}

	})

	facts_hide := widget.NewLabel("1. В Португалии ненастная погода является уважительной причиной для неявки на работу.\n2. В некоторых засушливых странах при приветствии люди желают друг другу не здоровья, а дождя.\n3. В 1932 году в США из-за сильного холода целиком замерз Ниагарский водопад.\n4. Мужчины имеют в 6 раз выше вероятность быть пораженными молнией, чем женщины.\n5. В каплях дождя содержится витамин В12.\n6. Одно из самых солнечных мест на земле — Мертвое море. Здесь в среднем бывает около 330 солнечных дней в году.\n7. Самое пасмурное место на земле — архипелаг Северная Земля. Тут солнце светит всего 12 дней в году.\n8. В пустыне Сахара был зафиксирован снег только однажды, 18 февраля 1979 года.\n9. За десять минут ураган высвобождает больше энергии, чем все ядерное оружие в мире.\n10. Каждый час на Земле происходит около 760 гроз.")
	facts_hide.TextStyle.Symbol = true
	facts_hide.Hidden = true
	button_Weather_facts := widget.NewButton("Немного фактов о погоде 🌍 (Показать/Скрыть)", func() {
		facts_hide.Hidden = !facts_hide.Hidden
	})

	line := canvas.NewLine(color.Transparent)

	label_poesy := canvas.NewText("Немного поэзии может успокоить мысли:", color_hi)
	label_poesy.TextStyle.Bold = true
	cont_poesy := container.NewCenter(
		label_poesy,
	)

	App_poesy := container.NewAppTabs(
		container.NewTabItem("О бессоннице", widget.NewLabel("Ко мне не придёт успокоющий сон,\nПодобный прохладной и свежей реке.\nНи разу им больше не буду спасён.\nЯ слышу его теперь лишь вдалеке...\n\nСон не погасит огней моих мыслей,\nНе остановит сплошной карусели,\nВ которую с режущим скрипом и свистом\nСмазались все мои дни и недели.\n\nСон не прольёт остужающим ливнем.\nЕму не по силам сразиться с пожаром,\nКоторый, плюясь чёрным дымом противным,\nПлавит металл головы моей ржавой.\n\nСон не омоет кровей под ногтями,\nТающей жизни холодных эссенций.\nИ я буду лезвия красить, пока не\nВзовёртся в груди неуспыпное сердце...")),
		container.NewTabItem("Про Яблочный Меч", widget.NewLabel("Кто молекулу влёт на слои рассечёт?\nКто порезом лизнёт между звёзд небосвод?\nКто остёр как химера и сладок как мёд,\nНа ресницы льющийся нитью?\n\nОн, дразнясь, одолеет хоть сотню преград\nСквозь доспехи из стали, таланта и клятв,\nИ обгонит самый бдительный взгляд\nНесравненной яблочной прытью...")),
		container.NewTabItem("Помрачнее", widget.NewLabel("Слышу немного ветра,\nВижу немного звёзд.\nВнимая воздушным ответам,\nЧитаю на небе вопрос.\n\nПоднявшись с кровати, иду\nИ только потом просыпаюсь.\nВнезапно от боли кричу...\nИ только потом уже ранюсь.\n\nЯ делаю всё вопреки\nТой логике, что вам привычна.\nВ воронку от камня сойдутся круги,\nГасительный жест зажжёт спичку.\n\nМолчать от избытка эмоций...\nИ чай ледяной в молоко добавлять.\nЯ восходящее Солнце...\nНа западе преданно буду встречать.\n\nСо мною хорошее всё пойдёт плохо,\nЗажившие раны взойдут на лице.\nПод реквием тусклой последней эпохи\nПроснусь я в самом конце.")),
		container.NewTabItem("Лирические", widget.NewLabel("О нетронутой лишней любви\nЗапоёт и затихнет вскоре\nТа Луна, что колышет море\nМановеньем своей брови.\n\n***\n\nПодсказки судьбы различив скрупулёзно,\nТы знаешь, что не доживёшь до седин.\nТебя каждую ночь ослепляют те звёзды,\nКоторые видишь лишь ты один.")),
	)

	sheeps := []string{
		"Летучая Джилл",
		"Ефросиняя",
		"Помидора II",
		"Чёрная из Пустоши",
		"Заплетучка",
		"Бе-бе-бестолочь",
		"Лана сель дель Рей",
		"Мини-ралка",
		"Кучеряга Старшая",
		"Гретта Ноги-Базуки",
	}

	bnt_sheeps := widget.NewButton("Да, разумеется, подойдут и старые-добрые способы.\n Порой здесь пробегают овечки, и даже устраивают скоростные забеги.\nМожно попытаться их сосчитать... Или угадать, кто нынче всех взгреет? Поехали!🐑", func() {
		s := a.NewWindow("Блуждающие овцы")
		s.Resize(fyne.NewSize(400, 200))
		s.CenterOnScreen()

		sheepOrAsk := widget.NewLabel("Гонка началась!")

		Sheeps_bar := widget.NewProgressBarInfinite()
		Sheeps_bar.Start()

		cont := container.NewVBox(
			Sheeps_bar,
			sheepOrAsk,
		)

		s.SetContent(cont)
		s.Show()

		sheep(sheepOrAsk, sheeps)
	})

	cont := container.NewVBox(
		cont_hi,
		sel,
		label_Weather,
		btn_Weather_see,
		button_Weather_facts,
		facts_hide,
		img_girl,
		line, // Для отступа
		cont_poesy,
		App_poesy,
		bnt_sheeps,
	)

	scr_cont := container.NewScroll(cont)
	w.SetContent(scr_cont)
	w.Show()
	a.Run()
}

func sheep(sheepOrAsk *widget.Label, sheeps []string) {
	ch := make(chan string)
	ch2 := make(chan string)
	wg := &sync.WaitGroup{}

	for range 3 {
		go func() {
			ch <- sheeps[rand.Intn(len(sheeps))]
		}()
	}

	go func() {
		ch2 <- sheeps[rand.Intn(len(sheeps))]
	}()

	go func() {
		wg.Add(3)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(2) * time.Second)
			fyne.Do(func() {
				sheepOrAsk.SetText(fmt.Sprintf("Вот это да!\nВперёд нежданно вырывается %v!\nЗа ней устремляется %v!", <-ch, <-ch2))
				sheepOrAsk.Refresh()
			})
		}()

		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(6) * time.Second)
			fyne.Do(func() {
				sheepOrAsk.SetText(fmt.Sprintf("Невероятно! У самого финиша %v чуть не вылетает с трассы!\nЧто творится?!", <-ch))
			})
		}()

		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(10) * time.Second)
			fyne.Do(func() {
				sheepOrAsk.SetText(fmt.Sprintf("И первой приходит: %v!\nОстальные глотают пыль... Аплодисменты!", <-ch))
			})
		}()

		go func() {
			time.Sleep(time.Duration(15) * time.Second)
			wg.Wait()
			close(ch)
			close(ch2)
		}()
	}()
}
