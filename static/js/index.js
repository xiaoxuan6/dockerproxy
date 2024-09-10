function append(data) {
    let innerHtml = ''
    $.each(data, function (index, re) {
        innerHtml += '<div class="w-full flex justify-center items-center">\n' +
            '                <div class="w-52 h-12 relative rounded overflow-hidden z-100 before:content-[&quot;&quot;] before:h-24 before:w-24 before:rounded-full before:absolute before:z-200 before:inset-0 before:left-0 before:top-0 before:bg-radial-gradient before:translate-x-[var(--x,10000px)] before:translate-y-[var(--y,10000px)]">\n' +
            '                    <div class="p-2 flex flex-row items-center justify-between space-x-2 cursor-pointer absolute inset-[1px] rounded z-300 bg-gray-100 dark:bg-black scroll-container">\n' +
            '                        <div class="max-w-5xl overflow-clip hover:overflow-auto">\n' +
            '                            <p class="dark:text-gray-200 transition-all duration-300 scroll-text text-nowrap">\n' +
            '                               <a href="' + re.url + '" target="_blank">' + re.url + '</a>' +
            '                            </p>\n' +
            '                        </div>\n' +
            '                    </div>\n' +
            '                    <div aria-valuemax="100" aria-valuemin="0" role="progressbar" data-state="indeterminate"\n' +
            '                         data-max="100"\n' +
            '                         class="h-1 rounded-full bg-indigo-500/20 w-[96%] absolute bottom-[3px] left-[2%] overflow-hidden">\n' +
            '                        <div data-state="indeterminate" data-max="100"\n';

        if (re.stat === true) {
            innerHtml += 'class="h-full w-full flex-1 bg-indigo-500 transition-all"\n';
        } else {
            innerHtml += 'class="h-full w-full flex-1 bg-red-500 transition-all"\n';
        }

        innerHtml += 'style="transform:translateX(-0%)"></div>\n' +
            '                    </div>\n' +
            '                </div>\n' +
            '            </div>'
    })

    document.getElementsByClassName('contents')[0].innerHTML = innerHtml;
}