const test = [
    [
        [0, 9],
        [5, 9],
    ],
    [
        [8, 0],
        [0, 8],
    ],
    [
        [9, 4],
        [3, 4],
    ],
    [
        [2, 2],
        [2, 1],
    ],
    [
        [7, 0],
        [7, 4],
    ],
    [
        [6, 4],
        [2, 0],
    ],
    [
        [0, 9],
        [2, 9],
    ],
    [
        [3, 4],
        [1, 4],
    ],
    [
        [0, 0],
        [8, 8],
    ],
    [
        [5, 5],
        [8, 2],
    ],
]

const input = [
    [
        [941, 230],
        [322, 849],
    ],
    [
        [762, 196],
        [701, 257],
    ],
    [
        [656, 197],
        [595, 136],
    ],
    [
        [687, 692],
        [57, 692],
    ],
    [
        [37, 953],
        [903, 87],
    ],
    [
        [674, 102],
        [84, 102],
    ],
    [
        [952, 323],
        [786, 157],
    ],
    [
        [807, 948],
        [430, 948],
    ],
    [
        [280, 66],
        [514, 66],
    ],
    [
        [810, 381],
        [928, 263],
    ],
    [
        [41, 278],
        [112, 207],
    ],
    [
        [754, 11],
        [754, 574],
    ],
    [
        [499, 830],
        [725, 604],
    ],
    [
        [713, 172],
        [658, 172],
    ],
    [
        [805, 54],
        [594, 54],
    ],
    [
        [442, 910],
        [40, 508],
    ],
    [
        [160, 170],
        [925, 935],
    ],
    [
        [265, 899],
        [265, 313],
    ],
    [
        [960, 976],
        [77, 93],
    ],
    [
        [820, 244],
        [877, 187],
    ],
    [
        [883, 501],
        [345, 501],
    ],
    [
        [12, 978],
        [941, 49],
    ],
    [
        [988, 46],
        [988, 572],
    ],
    [
        [285, 775],
        [285, 298],
    ],
    [
        [718, 69],
        [121, 69],
    ],
    [
        [218, 641],
        [146, 641],
    ],
    [
        [857, 277],
        [124, 277],
    ],
    [
        [32, 36],
        [657, 36],
    ],
    [
        [964, 280],
        [609, 280],
    ],
    [
        [739, 981],
        [910, 981],
    ],
    [
        [960, 794],
        [243, 794],
    ],
    [
        [447, 682],
        [751, 378],
    ],
    [
        [813, 103],
        [813, 240],
    ],
    [
        [568, 705],
        [497, 705],
    ],
    [
        [888, 47],
        [888, 231],
    ],
    [
        [936, 95],
        [336, 695],
    ],
    [
        [305, 349],
        [18, 636],
    ],
    [
        [54, 240],
        [54, 222],
    ],
    [
        [28, 704],
        [625, 107],
    ],
    [
        [680, 325],
        [680, 623],
    ],
    [
        [209, 405],
        [209, 123],
    ],
    [
        [947, 906],
        [947, 721],
    ],
    [
        [149, 810],
        [834, 125],
    ],
    [
        [897, 875],
        [146, 124],
    ],
    [
        [928, 267],
        [928, 484],
    ],
    [
        [871, 516],
        [871, 136],
    ],
    [
        [954, 725],
        [706, 725],
    ],
    [
        [680, 645],
        [958, 645],
    ],
    [
        [680, 326],
        [908, 326],
    ],
    [
        [173, 157],
        [890, 874],
    ],
    [
        [842, 802],
        [166, 126],
    ],
    [
        [750, 442],
        [270, 922],
    ],
    [
        [567, 891],
        [567, 784],
    ],
    [
        [374, 623],
        [374, 174],
    ],
    [
        [979, 725],
        [765, 511],
    ],
    [
        [336, 440],
        [82, 440],
    ],
    [
        [214, 213],
        [939, 938],
    ],
    [
        [652, 815],
        [763, 815],
    ],
    [
        [220, 48],
        [331, 159],
    ],
    [
        [580, 522],
        [141, 522],
    ],
    [
        [286, 685],
        [286, 779],
    ],
    [
        [865, 343],
        [865, 257],
    ],
    [
        [738, 898],
        [405, 565],
    ],
    [
        [703, 571],
        [420, 571],
    ],
    [
        [792, 368],
        [792, 955],
    ],
    [
        [738, 905],
        [738, 79],
    ],
    [
        [646, 95],
        [737, 95],
    ],
    [
        [930, 908],
        [72, 50],
    ],
    [
        [310, 933],
        [310, 243],
    ],
    [
        [192, 22],
        [918, 748],
    ],
    [
        [245, 803],
        [81, 639],
    ],
    [
        [567, 218],
        [901, 218],
    ],
    [
        [148, 950],
        [965, 133],
    ],
    [
        [147, 772],
        [159, 772],
    ],
    [
        [774, 84],
        [774, 960],
    ],
    [
        [860, 798],
        [372, 798],
    ],
    [
        [856, 131],
        [856, 703],
    ],
    [
        [368, 603],
        [247, 603],
    ],
    [
        [587, 533],
        [301, 533],
    ],
    [
        [832, 461],
        [832, 506],
    ],
    [
        [164, 709],
        [960, 709],
    ],
    [
        [874, 471],
        [327, 471],
    ],
    [
        [346, 237],
        [346, 921],
    ],
    [
        [683, 300],
        [910, 527],
    ],
    [
        [353, 717],
        [353, 575],
    ],
    [
        [586, 578],
        [798, 366],
    ],
    [
        [27, 813],
        [27, 434],
    ],
    [
        [311, 391],
        [418, 391],
    ],
    [
        [369, 304],
        [33, 304],
    ],
    [
        [591, 226],
        [591, 558],
    ],
    [
        [634, 545],
        [513, 545],
    ],
    [
        [439, 257],
        [207, 257],
    ],
    [
        [42, 791],
        [581, 252],
    ],
    [
        [155, 801],
        [155, 294],
    ],
    [
        [599, 603],
        [599, 182],
    ],
    [
        [48, 607],
        [337, 896],
    ],
    [
        [199, 828],
        [506, 828],
    ],
    [
        [28, 147],
        [733, 852],
    ],
    [
        [799, 563],
        [799, 22],
    ],
    [
        [206, 625],
        [455, 874],
    ],
    [
        [185, 330],
        [335, 480],
    ],
    [
        [161, 746],
        [590, 746],
    ],
    [
        [932, 13],
        [269, 13],
    ],
    [
        [649, 746],
        [649, 309],
    ],
    [
        [463, 169],
        [930, 636],
    ],
    [
        [568, 251],
        [386, 251],
    ],
    [
        [739, 692],
        [233, 692],
    ],
    [
        [941, 989],
        [84, 132],
    ],
    [
        [513, 356],
        [513, 628],
    ],
    [
        [534, 168],
        [285, 168],
    ],
    [
        [447, 563],
        [447, 698],
    ],
    [
        [898, 915],
        [791, 808],
    ],
    [
        [339, 405],
        [432, 405],
    ],
    [
        [414, 940],
        [335, 940],
    ],
    [
        [591, 741],
        [59, 741],
    ],
    [
        [347, 330],
        [347, 341],
    ],
    [
        [186, 40],
        [438, 292],
    ],
    [
        [849, 872],
        [295, 318],
    ],
    [
        [406, 620],
        [938, 620],
    ],
    [
        [346, 226],
        [864, 226],
    ],
    [
        [609, 40],
        [478, 171],
    ],
    [
        [820, 900],
        [947, 900],
    ],
    [
        [201, 63],
        [201, 107],
    ],
    [
        [984, 652],
        [47, 652],
    ],
    [
        [193, 204],
        [776, 204],
    ],
    [
        [173, 892],
        [740, 892],
    ],
    [
        [389, 675],
        [709, 355],
    ],
    [
        [489, 954],
        [546, 954],
    ],
    [
        [18, 82],
        [587, 651],
    ],
    [
        [646, 150],
        [675, 150],
    ],
    [
        [618, 805],
        [618, 592],
    ],
    [
        [178, 617],
        [178, 606],
    ],
    [
        [179, 30],
        [505, 30],
    ],
    [
        [984, 21],
        [21, 984],
    ],
    [
        [172, 167],
        [15, 167],
    ],
    [
        [17, 209],
        [192, 209],
    ],
    [
        [814, 945],
        [814, 18],
    ],
    [
        [385, 632],
        [161, 632],
    ],
    [
        [126, 41],
        [474, 389],
    ],
    [
        [575, 778],
        [737, 778],
    ],
    [
        [74, 270],
        [147, 270],
    ],
    [
        [891, 248],
        [467, 672],
    ],
    [
        [95, 426],
        [95, 728],
    ],
    [
        [235, 73],
        [235, 583],
    ],
    [
        [730, 302],
        [730, 466],
    ],
    [
        [388, 587],
        [377, 598],
    ],
    [
        [525, 155],
        [184, 155],
    ],
    [
        [370, 278],
        [966, 874],
    ],
    [
        [950, 150],
        [444, 656],
    ],
    [
        [644, 935],
        [401, 935],
    ],
    [
        [798, 515],
        [506, 807],
    ],
    [
        [976, 562],
        [253, 562],
    ],
    [
        [674, 350],
        [603, 421],
    ],
    [
        [686, 653],
        [576, 653],
    ],
    [
        [691, 278],
        [593, 180],
    ],
    [
        [964, 961],
        [76, 73],
    ],
    [
        [735, 582],
        [735, 389],
    ],
    [
        [786, 885],
        [76, 885],
    ],
    [
        [402, 732],
        [231, 732],
    ],
    [
        [660, 881],
        [660, 525],
    ],
    [
        [683, 383],
        [683, 364],
    ],
    [
        [174, 20],
        [174, 75],
    ],
    [
        [692, 819],
        [107, 819],
    ],
    [
        [344, 669],
        [577, 902],
    ],
    [
        [562, 126],
        [697, 261],
    ],
    [
        [621, 344],
        [621, 707],
    ],
    [
        [731, 892],
        [213, 374],
    ],
    [
        [216, 828],
        [663, 828],
    ],
    [
        [990, 534],
        [990, 356],
    ],
    [
        [973, 714],
        [519, 714],
    ],
    [
        [25, 981],
        [983, 23],
    ],
    [
        [659, 399],
        [535, 275],
    ],
    [
        [967, 885],
        [183, 101],
    ],
    [
        [612, 684],
        [732, 684],
    ],
    [
        [955, 485],
        [955, 806],
    ],
    [
        [582, 714],
        [582, 719],
    ],
    [
        [342, 203],
        [905, 203],
    ],
    [
        [188, 488],
        [272, 488],
    ],
    [
        [659, 65],
        [659, 679],
    ],
    [
        [306, 85],
        [605, 384],
    ],
    [
        [975, 847],
        [975, 353],
    ],
    [
        [742, 989],
        [742, 652],
    ],
    [
        [917, 524],
        [934, 524],
    ],
    [
        [890, 571],
        [662, 799],
    ],
    [
        [901, 791],
        [901, 118],
    ],
    [
        [631, 447],
        [114, 447],
    ],
    [
        [850, 28],
        [797, 28],
    ],
    [
        [842, 759],
        [91, 759],
    ],
    [
        [659, 538],
        [253, 944],
    ],
    [
        [693, 69],
        [693, 452],
    ],
    [
        [161, 515],
        [789, 515],
    ],
    [
        [892, 630],
        [892, 785],
    ],
    [
        [78, 947],
        [931, 947],
    ],
    [
        [561, 728],
        [11, 178],
    ],
    [
        [138, 842],
        [138, 133],
    ],
    [
        [890, 373],
        [628, 373],
    ],
    [
        [509, 370],
        [592, 370],
    ],
    [
        [982, 41],
        [185, 838],
    ],
    [
        [184, 210],
        [184, 218],
    ],
    [
        [390, 525],
        [390, 558],
    ],
    [
        [387, 151],
        [387, 39],
    ],
    [
        [718, 808],
        [833, 808],
    ],
    [
        [206, 234],
        [206, 620],
    ],
    [
        [84, 150],
        [84, 959],
    ],
    [
        [336, 468],
        [307, 468],
    ],
    [
        [764, 19],
        [739, 44],
    ],
    [
        [752, 607],
        [643, 607],
    ],
    [
        [233, 149],
        [112, 149],
    ],
    [
        [368, 612],
        [725, 255],
    ],
    [
        [929, 497],
        [909, 477],
    ],
    [
        [829, 274],
        [829, 190],
    ],
    [
        [312, 268],
        [312, 128],
    ],
    [
        [519, 18],
        [519, 552],
    ],
    [
        [896, 19],
        [140, 19],
    ],
    [
        [368, 727],
        [368, 114],
    ],
    [
        [233, 813],
        [750, 813],
    ],
    [
        [477, 758],
        [477, 213],
    ],
    [
        [615, 171],
        [615, 530],
    ],
    [
        [38, 461],
        [301, 461],
    ],
    [
        [862, 107],
        [154, 815],
    ],
    [
        [271, 52],
        [271, 517],
    ],
    [
        [203, 936],
        [365, 936],
    ],
    [
        [96, 700],
        [13, 617],
    ],
    [
        [290, 554],
        [389, 455],
    ],
    [
        [377, 923],
        [377, 890],
    ],
    [
        [347, 511],
        [147, 511],
    ],
    [
        [889, 412],
        [762, 412],
    ],
    [
        [558, 412],
        [424, 412],
    ],
    [
        [45, 838],
        [45, 845],
    ],
    [
        [958, 27],
        [958, 454],
    ],
    [
        [154, 244],
        [20, 244],
    ],
    [
        [315, 154],
        [315, 173],
    ],
    [
        [135, 618],
        [135, 71],
    ],
    [
        [380, 422],
        [131, 671],
    ],
    [
        [314, 500],
        [314, 873],
    ],
    [
        [915, 320],
        [915, 159],
    ],
    [
        [213, 772],
        [977, 772],
    ],
    [
        [14, 22],
        [978, 986],
    ],
    [
        [444, 759],
        [444, 385],
    ],
    [
        [730, 650],
        [730, 210],
    ],
    [
        [532, 551],
        [633, 652],
    ],
    [
        [547, 426],
        [335, 426],
    ],
    [
        [868, 191],
        [156, 903],
    ],
    [
        [462, 599],
        [611, 748],
    ],
    [
        [729, 709],
        [729, 714],
    ],
    [
        [665, 229],
        [849, 413],
    ],
    [
        [880, 947],
        [880, 159],
    ],
    [
        [249, 837],
        [249, 604],
    ],
    [
        [575, 205],
        [196, 584],
    ],
    [
        [960, 665],
        [320, 25],
    ],
    [
        [617, 853],
        [412, 853],
    ],
    [
        [224, 60],
        [224, 467],
    ],
    [
        [226, 741],
        [226, 47],
    ],
    [
        [371, 595],
        [118, 342],
    ],
    [
        [371, 708],
        [371, 561],
    ],
    [
        [236, 141],
        [955, 860],
    ],
    [
        [55, 509],
        [55, 938],
    ],
    [
        [684, 885],
        [684, 670],
    ],
    [
        [93, 509],
        [497, 105],
    ],
    [
        [284, 61],
        [812, 61],
    ],
    [
        [438, 353],
        [242, 353],
    ],
    [
        [77, 716],
        [363, 430],
    ],
    [
        [283, 769],
        [905, 147],
    ],
    [
        [56, 799],
        [551, 799],
    ],
    [
        [804, 637],
        [804, 526],
    ],
    [
        [476, 54],
        [154, 54],
    ],
    [
        [686, 400],
        [686, 145],
    ],
    [
        [740, 905],
        [417, 905],
    ],
    [
        [21, 113],
        [823, 915],
    ],
    [
        [286, 132],
        [880, 726],
    ],
    [
        [923, 378],
        [771, 378],
    ],
    [
        [924, 922],
        [36, 34],
    ],
    [
        [801, 609],
        [801, 407],
    ],
    [
        [465, 671],
        [550, 756],
    ],
    [
        [628, 235],
        [628, 842],
    ],
    [
        [684, 840],
        [716, 808],
    ],
    [
        [841, 366],
        [495, 712],
    ],
    [
        [740, 208],
        [740, 174],
    ],
    [
        [657, 370],
        [657, 731],
    ],
    [
        [817, 781],
        [466, 781],
    ],
    [
        [308, 894],
        [308, 370],
    ],
    [
        [497, 233],
        [755, 233],
    ],
    [
        [35, 145],
        [35, 398],
    ],
    [
        [383, 163],
        [578, 163],
    ],
    [
        [620, 985],
        [620, 849],
    ],
    [
        [178, 253],
        [178, 724],
    ],
    [
        [556, 51],
        [556, 525],
    ],
    [
        [650, 187],
        [706, 243],
    ],
    [
        [161, 988],
        [599, 550],
    ],
    [
        [861, 256],
        [501, 616],
    ],
    [
        [46, 555],
        [181, 555],
    ],
    [
        [980, 975],
        [980, 916],
    ],
    [
        [345, 751],
        [479, 617],
    ],
    [
        [534, 642],
        [534, 202],
    ],
    [
        [901, 240],
        [901, 490],
    ],
    [
        [984, 280],
        [337, 927],
    ],
    [
        [578, 663],
        [578, 298],
    ],
    [
        [377, 943],
        [259, 943],
    ],
    [
        [975, 38],
        [39, 974],
    ],
    [
        [697, 870],
        [387, 560],
    ],
    [
        [147, 520],
        [218, 520],
    ],
    [
        [683, 711],
        [486, 711],
    ],
    [
        [825, 26],
        [122, 729],
    ],
    [
        [855, 84],
        [751, 84],
    ],
    [
        [558, 945],
        [989, 945],
    ],
    [
        [660, 195],
        [597, 195],
    ],
    [
        [889, 696],
        [317, 696],
    ],
    [
        [969, 248],
        [240, 977],
    ],
    [
        [598, 625],
        [598, 148],
    ],
    [
        [176, 151],
        [256, 151],
    ],
    [
        [939, 70],
        [648, 70],
    ],
    [
        [645, 431],
        [411, 431],
    ],
    [
        [502, 518],
        [221, 518],
    ],
    [
        [821, 988],
        [213, 988],
    ],
    [
        [361, 850],
        [684, 850],
    ],
    [
        [506, 173],
        [506, 405],
    ],
    [
        [323, 151],
        [726, 151],
    ],
    [
        [131, 519],
        [35, 519],
    ],
    [
        [164, 445],
        [798, 445],
    ],
    [
        [425, 989],
        [425, 133],
    ],
    [
        [18, 739],
        [684, 73],
    ],
    [
        [138, 545],
        [138, 155],
    ],
    [
        [401, 104],
        [766, 104],
    ],
    [
        [864, 855],
        [203, 855],
    ],
    [
        [636, 361],
        [604, 361],
    ],
    [
        [820, 970],
        [820, 882],
    ],
    [
        [866, 859],
        [835, 859],
    ],
    [
        [112, 507],
        [112, 715],
    ],
    [
        [529, 494],
        [529, 928],
    ],
    [
        [104, 469],
        [193, 469],
    ],
    [
        [82, 841],
        [831, 92],
    ],
    [
        [258, 518],
        [258, 778],
    ],
    [
        [34, 917],
        [135, 917],
    ],
    [
        [777, 553],
        [985, 345],
    ],
    [
        [64, 952],
        [719, 297],
    ],
    [
        [341, 224],
        [902, 224],
    ],
    [
        [87, 128],
        [525, 566],
    ],
    [
        [951, 400],
        [448, 903],
    ],
    [
        [344, 963],
        [21, 963],
    ],
    [
        [983, 244],
        [983, 503],
    ],
    [
        [938, 771],
        [635, 771],
    ],
    [
        [560, 262],
        [560, 974],
    ],
    [
        [46, 386],
        [75, 386],
    ],
    [
        [898, 747],
        [898, 17],
    ],
    [
        [239, 929],
        [149, 929],
    ],
    [
        [849, 881],
        [849, 251],
    ],
    [
        [204, 204],
        [204, 753],
    ],
    [
        [830, 33],
        [830, 130],
    ],
    [
        [304, 339],
        [42, 339],
    ],
    [
        [565, 312],
        [773, 312],
    ],
    [
        [387, 523],
        [234, 523],
    ],
    [
        [239, 421],
        [543, 725],
    ],
    [
        [197, 433],
        [197, 723],
    ],
    [
        [595, 21],
        [370, 21],
    ],
    [
        [547, 171],
        [480, 104],
    ],
    [
        [639, 910],
        [639, 241],
    ],
    [
        [908, 185],
        [560, 185],
    ],
    [
        [947, 565],
        [947, 411],
    ],
    [
        [211, 670],
        [588, 293],
    ],
    [
        [753, 708],
        [753, 624],
    ],
    [
        [36, 147],
        [859, 970],
    ],
    [
        [423, 94],
        [930, 94],
    ],
    [
        [613, 680],
        [607, 680],
    ],
    [
        [277, 263],
        [836, 822],
    ],
    [
        [186, 413],
        [827, 413],
    ],
    [
        [483, 173],
        [142, 173],
    ],
    [
        [25, 771],
        [409, 387],
    ],
    [
        [328, 916],
        [613, 631],
    ],
    [
        [267, 604],
        [724, 147],
    ],
    [
        [430, 616],
        [150, 896],
    ],
    [
        [692, 463],
        [50, 463],
    ],
    [
        [306, 360],
        [306, 653],
    ],
    [
        [736, 948],
        [736, 174],
    ],
    [
        [797, 529],
        [774, 529],
    ],
    [
        [492, 486],
        [492, 812],
    ],
    [
        [659, 429],
        [102, 429],
    ],
    [
        [582, 503],
        [695, 616],
    ],
    [
        [780, 62],
        [780, 164],
    ],
    [
        [58, 318],
        [387, 318],
    ],
    [
        [286, 694],
        [286, 396],
    ],
    [
        [248, 241],
        [248, 361],
    ],
    [
        [112, 963],
        [707, 963],
    ],
    [
        [771, 722],
        [636, 722],
    ],
    [
        [508, 76],
        [389, 76],
    ],
    [
        [435, 307],
        [201, 541],
    ],
    [
        [167, 312],
        [618, 763],
    ],
    [
        [721, 407],
        [305, 823],
    ],
    [
        [57, 203],
        [516, 203],
    ],
    [
        [83, 239],
        [83, 607],
    ],
    [
        [810, 686],
        [137, 13],
    ],
    [
        [817, 268],
        [101, 984],
    ],
    [
        [379, 975],
        [379, 631],
    ],
    [
        [597, 38],
        [611, 38],
    ],
    [
        [56, 504],
        [56, 900],
    ],
    [
        [108, 587],
        [261, 740],
    ],
    [
        [625, 426],
        [476, 426],
    ],
    [
        [248, 486],
        [643, 881],
    ],
    [
        [932, 25],
        [21, 936],
    ],
    [
        [388, 613],
        [388, 296],
    ],
    [
        [644, 188],
        [644, 273],
    ],
    [
        [871, 425],
        [871, 791],
    ],
    [
        [722, 866],
        [722, 39],
    ],
    [
        [96, 579],
        [96, 97],
    ],
    [
        [876, 64],
        [297, 643],
    ],
    [
        [581, 633],
        [59, 633],
    ],
    [
        [11, 10],
        [989, 988],
    ],
    [
        [947, 55],
        [266, 736],
    ],
    [
        [532, 553],
        [735, 756],
    ],
    [
        [898, 855],
        [83, 40],
    ],
    [
        [533, 289],
        [306, 62],
    ],
    [
        [497, 736],
        [332, 571],
    ],
    [
        [871, 201],
        [345, 727],
    ],
    [
        [550, 686],
        [256, 686],
    ],
    [
        [858, 585],
        [607, 836],
    ],
    [
        [380, 171],
        [15, 171],
    ],
    [
        [864, 112],
        [864, 686],
    ],
    [
        [791, 857],
        [305, 857],
    ],
    [
        [898, 579],
        [741, 579],
    ],
    [
        [479, 713],
        [113, 713],
    ],
    [
        [19, 143],
        [779, 903],
    ],
    [
        [347, 161],
        [140, 368],
    ],
    [
        [479, 395],
        [534, 340],
    ],
    [
        [929, 37],
        [77, 889],
    ],
    [
        [128, 958],
        [884, 202],
    ],
    [
        [921, 18],
        [921, 650],
    ],
    [
        [263, 550],
        [263, 280],
    ],
    [
        [155, 592],
        [235, 592],
    ],
    [
        [565, 34],
        [565, 454],
    ],
    [
        [913, 371],
        [173, 371],
    ],
    [
        [199, 158],
        [974, 933],
    ],
    [
        [98, 775],
        [98, 234],
    ],
    [
        [649, 576],
        [649, 444],
    ],
    [
        [801, 855],
        [548, 855],
    ],
    [
        [859, 913],
        [363, 913],
    ],
    [
        [274, 487],
        [274, 654],
    ],
    [
        [729, 982],
        [443, 982],
    ],
    [
        [664, 827],
        [77, 240],
    ],
    [
        [656, 885],
        [656, 350],
    ],
    [
        [916, 74],
        [284, 706],
    ],
    [
        [439, 31],
        [439, 175],
    ],
    [
        [423, 753],
        [280, 753],
    ],
    [
        [424, 914],
        [948, 914],
    ],
    [
        [980, 723],
        [980, 674],
    ],
    [
        [656, 437],
        [626, 407],
    ],
    [
        [577, 654],
        [423, 654],
    ],
    [
        [19, 224],
        [424, 224],
    ],
    [
        [310, 181],
        [704, 575],
    ],
    [
        [828, 296],
        [828, 308],
    ],
    [
        [905, 151],
        [955, 151],
    ],
    [
        [319, 178],
        [892, 178],
    ],
    [
        [972, 939],
        [65, 32],
    ],
    [
        [497, 98],
        [91, 98],
    ],
    [
        [987, 402],
        [943, 446],
    ],
    [
        [904, 19],
        [174, 749],
    ],
    [
        [265, 885],
        [265, 835],
    ],
    [
        [475, 414],
        [658, 597],
    ],
    [
        [610, 93],
        [938, 93],
    ],
    [
        [961, 892],
        [661, 892],
    ],
    [
        [297, 600],
        [378, 600],
    ],
    [
        [405, 637],
        [52, 284],
    ],
    [
        [439, 874],
        [439, 612],
    ],
    [
        [275, 185],
        [275, 218],
    ],
    [
        [220, 840],
        [220, 735],
    ],
    [
        [372, 153],
        [644, 425],
    ],
    [
        [896, 964],
        [896, 461],
    ],
    [
        [916, 484],
        [951, 449],
    ],
    [
        [485, 355],
        [456, 355],
    ],
    [
        [198, 793],
        [198, 132],
    ],
    [
        [614, 735],
        [561, 735],
    ],
    [
        [181, 591],
        [147, 591],
    ],
    [
        [175, 289],
        [159, 289],
    ],
    [
        [899, 758],
        [962, 695],
    ],
    [
        [506, 647],
        [506, 858],
    ],
    [
        [443, 828],
        [720, 828],
    ],
    [
        [623, 641],
        [623, 631],
    ],
    [
        [202, 409],
        [891, 409],
    ],
    [
        [486, 751],
        [80, 345],
    ],
    [
        [781, 73],
        [781, 710],
    ],
    [
        [911, 643],
        [911, 571],
    ],
    [
        [799, 151],
        [89, 861],
    ],
    [
        [716, 815],
        [810, 815],
    ],
    [
        [947, 517],
        [947, 575],
    ],
    [
        [704, 260],
        [704, 727],
    ],
    [
        [113, 581],
        [113, 606],
    ],
    [
        [408, 252],
        [408, 761],
    ],
    [
        [601, 753],
        [457, 609],
    ],
    [
        [851, 424],
        [501, 774],
    ],
    [
        [670, 941],
        [916, 941],
    ],
    [
        [480, 839],
        [205, 564],
    ],
    [
        [912, 949],
        [38, 75],
    ],
    [
        [477, 39],
        [925, 487],
    ],
    [
        [139, 898],
        [309, 898],
    ],
    [
        [93, 386],
        [93, 194],
    ],
    [
        [184, 132],
        [943, 891],
    ],
    [
        [247, 557],
        [247, 182],
    ],
    [
        [832, 22],
        [76, 778],
    ],
    [
        [61, 814],
        [806, 69],
    ],
    [
        [816, 640],
        [604, 428],
    ],
    [
        [214, 561],
        [623, 152],
    ],
    [
        [698, 858],
        [389, 858],
    ],
]

// const getValuesBetween = (num1, num2) => [...Array(Math.abs(num1 - num2) + 1).keys()].map((i) => i + Math.min(num1, num2))
// const getPointsOnLine = ([[x1, y1], [x2, y2]]) =>
//     getValuesBetween(x1, x2)
//         .map((x) => (x1 != x2 && y1 != y2 ? [] : getValuesBetween(y1, y2).map((y) => `${x}${y}`)))
//         .flat()

// const getPointsOnLines = (lines) => lines.map((line) => getPointsOnLine(line)).flat()

// const markPoints = (points) =>
//     points.reduce((prev, point) => {
//         const found = prev.find((i) => i.key === point)
//         if (found) found.value++
//         return found ? prev : [...prev, { key: point, value: 1 }]
//     }, [])

// console.log(markPoints(getPointsOnLines(input)).filter((point) => point.value >= 2).length)

const maxXY = (lines) => {
    const maxPerLine = lines.map(([[x1, y1], [x2, y2]]) => [Math.max(x1, x2), Math.max(y1, y2)])
    const maxX = Math.max(...maxPerLine.map((line) => line[0]))
    const maxY = Math.max(...maxPerLine.map((line) => line[1]))
    return [maxX, maxY]
}

const filterDiagonals = (lines) => lines.filter(([[x1, y1], [x2, y2]]) => x1 === x2 || y1 === y2)
const getGrid = ([maxX, maxY]) => [...Array(maxX + 1)].map((_) => [...Array(maxY + 1)].map((_) => 0))
const isPointOnLine = ([x, y], [[x1, y1], [x2, y2]]) => {
    const isEndPoint = (x === x1 && y === y1) || (x === x2 && y === y2)
    if (isEndPoint) return true
    const isPointBetweenEnds = ((x1 <= x && x <= x2) || (x2 <= x && x <= x1)) && ((y1 <= y && y <= y2) || (y2 <= y && y <= y1))
    const isOrthogonal = x1 === x2 || y1 === y2
    if (isOrthogonal) return isPointBetweenEnds
    const sameSlope = Math.abs(y1 - y2) / Math.abs(x1 - x2) === Math.abs(y1 - y) / Math.abs(x1 - x)
    return isPointBetweenEnds && sameSlope
}
const countLinesWithPoint = (point, lines) => lines.map((line) => isPointOnLine(point, line)).filter((x) => x).length

console.log(
    getGrid(maxXY(input))
        .map((row, y) => {
            console.log(y)
            return row.map((_, x) => countLinesWithPoint([x, y], input))
        })
        .flat()
        .filter((x) => x >= 2).length
)
